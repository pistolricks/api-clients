import { createSignal, onMount, onCleanup, Show, createEffect, Component } from 'solid-js';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';
import OSM from 'ol/source/OSM';
import Feature from 'ol/Feature';
import Point from 'ol/geom/Point';
import { fromLonLat } from 'ol/proj';
import { Style, Circle, Fill, Stroke, Text } from 'ol/style';
import Overlay from 'ol/Overlay';
import { School, SchoolsResponse } from '~/types';
import 'ol/ol.css';
import {DrawerContent} from "~/components/ui/drawer/drawer";


const MapComponent: Component = () => {
    const [schools, setSchools] = createSignal<School[]>([]);
    const [loading, setLoading] = createSignal<boolean>(true);
    const [error, setError] = createSignal<string | null>(null);
    const [page, setPage] = createSignal<number>(1);
    const [pageSize, setPageSize] = createSignal<number>(100);
    const [total, setTotal] = createSignal<number>(0);
    const [selectedSchool, setSelectedSchool] = createSignal<School | null>(null);

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
                    fill: new Fill({ color: '#007bff' }),
                    stroke: new Stroke({ color: '#ffffff', width: 2 })
                }),
                text: new Text({
                    text: school.name,
                    offsetY: -15,
                    font: '12px Calibri,sans-serif',
                    fill: new Fill({ color: '#000' }),
                    stroke: new Stroke({ color: '#fff', width: 3 })
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

    return (
        <div class="app-container">
            <div class="header">
                <Show when={error()}>
                    <div class="error">{error()}</div>
                </Show>
            </div>

            <DrawerContent contextId={'map-drawer'}>
                TEST
            </DrawerContent>

            <div class="map-container ">
                <div ref={mapElement} class="ol-map"></div>
                <div ref={popupElement} class="school-popup" style={{ display: selectedSchool() ? 'block' : 'none' }}>
                    <Show when={selectedSchool()}>
                        <h3>{selectedSchool()!.name}</h3>
                        <p><strong>Address:</strong> {selectedSchool()!.address}, {selectedSchool()!.city}, {selectedSchool()!.state} {selectedSchool()!.zip}</p>
                        <p><strong>Type:</strong> {selectedSchool()!.level || 'N/A'}</p>
                        <p><strong>Enrollment:</strong> {selectedSchool()!.enrollment || 'N/A'}</p>
                        <p><strong>Coordinates:</strong> {selectedSchool()!.latitude.toFixed(6)}, {selectedSchool()!.longitude.toFixed(6)}</p>
                        <p>{selectedSchool()!.id}</p>

                        <p>{selectedSchool()!.country}</p>
                        <p>{selectedSchool()!.county}</p>
                        <p>{selectedSchool()!.countyfips}</p>
                        <p>{selectedSchool()!.latitude}</p>
                        <p>{selectedSchool()!.longitude}</p>
                        <p>{selectedSchool()!.level}</p>
                        <p>{selectedSchool()!.st_grade}</p>
                        <p>{selectedSchool()!.end_grade}</p>
                        <p>{selectedSchool()!.enrollment}</p>
                        <p>{selectedSchool()!.ft_teacher}</p>
                        <p>{selectedSchool()!.type}</p>
                        <p>{selectedSchool()!.status}</p>
                        <p>{selectedSchool()!.population}</p>
                        <p>{selectedSchool()!.ncesid}</p>
                        <p>{selectedSchool()!.districtid}</p>
                        <p>{selectedSchool()!.naics_code}</p>
                        <p>{selectedSchool()!.naics_desc}</p>
                        <p>{selectedSchool()!.website}</p>
                        <p>{selectedSchool()!.telephone}</p>
                        <p>{selectedSchool()!.sourcedate}</p>
                        <p>{selectedSchool()!.val_date}</p>
                        <p>{selectedSchool()!.val_method}</p>
                        <p>{selectedSchool()!.source}</p>
                        <p>{selectedSchool()!.shelter_id}</p>
                        <p>{selectedSchool()!.created_at}</p>
                        <p>{selectedSchool()!.updated_at}</p>
                    </Show>
                </div>
            </div>

            <div class="pagination absolute bottom-0 left-0 right-0">
                <button onClick={handlePrevPage} disabled={page() === 1 || loading()}>Previous</button>
                <span>Page {page()} of {Math.ceil(total() / pageSize())}</span>
                <button onClick={handleNextPage} disabled={page() * pageSize() >= total() || loading()}>Next</button>
            </div>
        </div>
    );
};

export default MapComponent;