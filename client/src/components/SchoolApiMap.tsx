import {Component, createEffect, createSignal, onMount, Show} from 'solid-js';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import OSM from 'ol/source/OSM';
import Feature from 'ol/Feature';
import Point from 'ol/geom/Point';
import {fromLonLat} from 'ol/proj';
import {Circle, Fill, Stroke, Style, Text} from 'ol/style';
import Overlay from 'ol/Overlay';
import {School, SchoolsResponse} from '~/types';
import 'ol/ol.css';
import BaseDrawer, {DrawerContent} from "~/components/ui/drawer/drawer";
import {SchoolDetails} from "~/components/schools/details";
import Icon from "~/components/ui/icon";


const MapComponent: Component = () => {
    const [schools, setSchools] = createSignal<School[]>([]);
    const [loading, setLoading] = createSignal<boolean>(true);
    const [error, setError] = createSignal<string | null>(null);
    const [page, setPage] = createSignal<number>(1);
    const [pageSize, setPageSize] = createSignal<number>(100);
    const [total, setTotal] = createSignal<number>(0);
    const [selectedSchool, setSelectedSchool] = createSignal<School | null>(null);
    const [getOpen, setOpen] = createSignal(false);
    let mapElement: HTMLDivElement | undefined;
    let popupElement: HTMLDivElement | undefined;
    let map: Map | undefined;
    let vectorSource: VectorSource | undefined;
    let popup: Overlay | undefined;

    const fetchSchools = async (): Promise<void> => {
        setLoading(true);
        setError(null);
        try {
            const response = await fetch(`/api/schools?page=${page()}&pageSize=${pageSize()}`);
            if (!response.ok) {
                throw new Error(`HTTP error! status: ${response.status}`);
            }
            const data: SchoolsResponse = await response.json();
            setSchools(data.schools);
            setTotal(data.total);
            updateMap(data.schools);
        } catch (err) {
            if (err instanceof Error) {
                setError(err.message);
                console.error('Error fetching schools:', err);
            } else {
                setError('An unknown error occurred');
                console.error('Unknown error:', err);
            }
        } finally {
            setLoading(false);
        }
    };

    const updateMap = (schools: School[]): void => {
        if (!map || !vectorSource) return;

        // Clear existing features
        vectorSource.clear();

        // Add new features
        const features = schools.map(school => {
            const feature = new Feature({
                geometry: new Point(fromLonLat([school.longitude, school.latitude])),
                school: school
            });

            feature.setStyle(new Style({
                image: new Circle({
                    radius: 6,
                    fill: new Fill({color: '#007bff'}),
                    stroke: new Stroke({color: '#ffffff', width: 2})
                }),
                text: new Text({
                    text: school.name,
                    offsetY: -15,
                    font: '12px Calibri,sans-serif',
                    fill: new Fill({color: '#000'}),
                    stroke: new Stroke({color: '#fff', width: 3})
                })
            }));

            return feature;
        });

        vectorSource.addFeatures(features);

        // Fit view to features if there are any
        if (features.length > 0) {
            map.getView().fit(vectorSource.getExtent(), {
                padding: [50, 50, 50, 50],
                maxZoom: 10
            });
        }
    };

    const handlePrevPage = (): void => {
        if (page() > 1) {
            setPage(page() - 1);
        }
    };

    const handleNextPage = (): void => {
        if (page() * pageSize() < total()) {
            setPage(page() + 1);
        }
    };

    const showSchoolPopup = (school: School, coordinate: number[]): void => {
        setSelectedSchool(school);
        setOpen(true);
        if (popup) {
            popup.setPosition(coordinate);

        }
    };

    onMount(() => {
        // Initialize vector source and layer
        vectorSource = new VectorSource();
        const vectorLayer = new VectorLayer({
            source: vectorSource
        });

        // Initialize map
        if (mapElement) {
            map = new Map({
                target: mapElement,
                layers: [
                    new TileLayer({
                        source: new OSM()
                    }),
                    vectorLayer
                ],
                view: new View({
                    center: fromLonLat([-98.5795, 39.8283]), // Center of USA
                    zoom: 4
                })
            });

            // Initialize popup overlay
            if (popupElement) {
                popup = new Overlay({
                    element: popupElement,
                    positioning: 'bottom-center',
                    stopEvent: false,
                    offset: [0, -10]
                });
                map.addOverlay(popup);

                // Add click handler to show popup
                map.on('click', (evt) => {
                    if (!map) return;
                    const feature = map.forEachFeatureAtPixel(evt.pixel, (feature) => feature);
                    if (feature) {
                        const school = feature.get('school') as School;
                        const geometry = feature.getGeometry();
                        if (geometry && geometry instanceof Point) {
                            showSchoolPopup(school, geometry.getCoordinates());
                        }
                    } else {
                        setSelectedSchool(null);

                        if (popup) {
                            popup.setPosition(undefined);
                            setOpen(false);
                        }
                    }
                });
            }
        }

        // Fetch schools
        fetchSchools();
    });

    // Watch for page changes
    onMount(() => {
        const pageWatcher = (): void => {
            fetchSchools();
        };


        // Create a computed signal that depends on page
        createEffect(() => {
            page();
            pageWatcher();


        });

        // No need for manual disposal as onCleanup will handle it automatically
    });


    createEffect(() => {
        console.log("open", getOpen())
    });

    return (
        <div class="app-container">
            <div class="header">
                <Show when={error()}>
                    <div class="error">{error()}</div>
                </Show>
            </div>
            <BaseDrawer open={getOpen()} setOpen={setOpen} side={"right"} contextId={'map-drawer'}>
                <DrawerContent class={''} side="right" contextId={'map-drawer'}>
                    {selectedSchool() && <SchoolDetails {...selectedSchool()!} />}
                </DrawerContent>

                <div class="map-container">
                    <div ref={mapElement} class="ol-map"></div>
                    <div ref={popupElement} class="select-none school-popup flex flex-col" style={{display: selectedSchool() ? 'block' : 'none'}}>
                        <Show when={selectedSchool()}>
                            <div class="col-span-1 divide-y divide-gray-200 rounded-lg bg-white shadow">
                                <div class="flex w-full items-center justify-between space-x-6 p-6">
                                    <div class="flex-1 truncate">
                                        <div class="flex items-center space-x-3">
                                            <h3 class="truncate text-sm font-medium text-gray-900">{selectedSchool()!.name}</h3>

                                        </div>
                                        <p class="mt-1 truncate text-sm text-gray-500">{selectedSchool()!.address}</p>
                                        <p class="text-sm capitalize leading-6 tracking-wide">{selectedSchool()!.city!.toLowerCase()}, {selectedSchool()!.state}</p>

                                    </div>

                                </div>
                                <div>
                                    <div class="-mt-px flex divide-x divide-gray-200">
                                        <div class="flex w-0 flex-1">
                                            <a target={"_blank"} href={selectedSchool()!.website}
                                               class="relative -mr-px inline-flex w-0 flex-1 items-center justify-center gap-x-3 rounded-bl-lg border border-transparent py-4 text-sm font-semibold text-gray-900">
                                                <Icon name={"Globe"} class="size-5 text-gray-400" />

                                            </a>
                                        </div>
                                        <div class="flex w-0 flex-1">
                                            <a target={"_blank"} href={selectedSchool()!.source}
                                               class="relative -mr-px inline-flex w-0 flex-1 items-center justify-center gap-x-3 rounded-bl-lg border border-transparent py-4 text-sm font-semibold text-gray-900">
                                                <Icon name={"Info"} class="size-5 text-gray-400" />

                                            </a>
                                        </div>
                                        <div class="-ml-px flex w-0 flex-1">
                                            <a href={`tel:+1-${selectedSchool()!.telephone}`}
                                               class="relative inline-flex w-0 flex-1 items-center justify-center gap-x-3 rounded-br-lg border border-transparent py-4 text-sm font-semibold text-gray-900">
                                                <svg class="size-5 text-gray-400" viewBox="0 0 20 20"
                                                     fill="currentColor" aria-hidden="true" data-slot="icon">
                                                    <path fill-rule="evenodd"
                                                          d="M2 3.5A1.5 1.5 0 0 1 3.5 2h1.148a1.5 1.5 0 0 1 1.465 1.175l.716 3.223a1.5 1.5 0 0 1-1.052 1.767l-.933.267c-.41.117-.643.555-.48.95a11.542 11.542 0 0 0 6.254 6.254c.395.163.833-.07.95-.48l.267-.933a1.5 1.5 0 0 1 1.767-1.052l3.223.716A1.5 1.5 0 0 1 18 15.352V16.5a1.5 1.5 0 0 1-1.5 1.5H15c-1.149 0-2.263-.15-3.326-.43A13.022 13.022 0 0 1 2.43 8.326 13.019 13.019 0 0 1 2 5V3.5Z"
                                                          clip-rule="evenodd"/>
                                                </svg>

                                            </a>
                                        </div>
                                    </div>
                                </div>
                            </div>



                        </Show>
                    </div>
                </div>

                <div class="pagination absolute bottom-0 left-0 right-0">
                    <button onClick={handlePrevPage} disabled={page() === 1 || loading()}>Previous</button>
                    <span>Page {page()} of {Math.ceil(total() / pageSize())}</span>
                    <button onClick={handleNextPage} disabled={page() * pageSize() >= total() || loading()}>Next
                    </button>
                </div>
            </BaseDrawer>
        </div>
    );
};

export default MapComponent;