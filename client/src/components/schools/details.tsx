import {Component} from "solid-js";
import {School} from "~/types";
import Icon from "~/components/ui/icon"
const SchoolDetails: Component<School> = props => {

    const id = () => props.id;
    const objectid = () => props.objectid;
    const level = () => props.level ?? "";
    const name = () => props.name.replace(level(), "").replace("SCHOOL", "");
    const address = () => props.address ?? "";
    const city = () => props.city ?? "";
    const state = () => props.state ?? "";
    const zip = () => props.zip ?? "";
    const country = () => props.country ?? "";
    const county = () => props.county ?? "";
    const countyfips = () => props.countyfips ?? "";
    const latitude = () => props.latitude;
    const longitude = () => props.longitude;

    const st_grade = () => props.st_grade ?? "";
    const end_grade = () => props.end_grade ?? "";
    const enrollment = () => props.enrollment ?? "NOT AVAILABLE";
    const ft_teacher = () => props.ft_teacher ?? "NOT AVAILABLE";
    const type = () => props.type ?? "";
    const status = () => props.status ?? "NOT AVAILABLE";
    const population = () => props.population ?? "NOT AVAILABLE";
    const ncesid = () => props.ncesid ?? "";
    const districtid = () => props.districtid ?? "";
    const naics_code = () => props.naics_code ?? "";
    const naics_desc = () => props.naics_desc ?? "";
    const website = () => props.website ?? "";
    const telephone = () => props.telephone ?? "";
    const sourcedate = () => props.sourcedate ?? "";
    const val_date = () => props.val_date ?? "";
    const val_method = () => props.val_method ?? "";
    const source = () => props.source ?? "";
    const shelter_id = () => props.shelter_id ?? "";
    const created_at = () => props.created_at;
    const updated_at = () => props.updated_at;

    const grades: { [key: string]: string } = {
        PK: "Pre-K",
        KG: "Kindergarten",
        "01": "First Grade",
        "02": "Second Grade",
        "03": "Third Grade",
        "04": "Fourth Grade",
        "05": "Fifth Grade",
        "06": "Sixth Grade",
        "07": "Seventh Grade",
        "08": "Eighth Grade",
        "09": "Freshman Year",
        "10": "Sophomore Year",
        "11": "Junior Year",
        "12": "Senior Year",

    }

    return (

        <div class="h-[100dvh] w-full overflow-y-auto">

                <div class="px-4 py-3 flex justify-between bg-gray-200">
                    <dt class="text-xs font-medium text-gray-900 flex justify-start items-center"> <span>        {naics_desc()}</span> </dt>
                    <dt class="text-xs font-medium text-gray-900">{objectid()}-{st_grade()}-{end_grade()} </dt>


                    {/*
                    <Icon name={"School"} class={"size-4"}/>
                    <dt class="text-xs font-medium text-gray-900">{countyfips()}-{objectid()}</dt>
                    */}
                </div>

            <div class="overflow-hidden bg-white shadow">

                <div class="border-t border-gray-100">

                    <div class="flex justify-start items-end px-4 py-3 space-x-4">

                        <div class="text-left col-span-1">
                            <h1 class="text-sm font-bold text-gray-900 text-pretty">{name()}</h1>
                            <p class="text-sm capitalize leading-6 tracking-wide">{address().toLowerCase()},</p>
                            <p class="text-sm capitalize leading-6 tracking-wide">{city().toLowerCase()}, {state()} {zip()}</p>
                            <div class="text-sm capitalize leading-6 tracking-wide">{county().toLowerCase()}, {country()}</div>
                        </div>
                    </div>
                    <dl class="divide-y divide-gray-100 border-t border-gray-100">
                        <div class="px-2">
                            <div class=" text-sm text-gray-900 ">
                                <ul role="list" class="divide-y divide-gray-100">
                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Phone"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="truncate font-medium select-all">{telephone()}</span>
                                                <span class="shrink-0 text-gray-400">Primary</span>
                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <a href="#"
                                               class="font-medium text-blue-600 hover:text-blue-500">CALL</a>
                                        </div>
                                    </li>
                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Globe"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span
                                                    class="font-medium">Website</span>
                                                <span class="shrink-0 text-gray-400 w-32 truncate">{website()}</span>
                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <a href={website()}
                                               target="_blank"
                                               class="font-medium text-blue-600 hover:text-blue-500">LINK</a>
                                        </div>
                                    </li>
                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Info"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span
                                                    class="font-medium">Source</span>
                                                <span class="shrink-0 text-gray-400 w-32 truncate">{source()}</span>
                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <a href={source()}
                                               target="_blank"
                                               class="font-medium text-blue-600 hover:text-blue-500">LINK</a>
                                        </div>
                                    </li>


                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"MapPin"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">
                                                    Latitude
                                                </span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{latitude()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"MapPin"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">
                                                    Longitude
                                                </span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{longitude()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Power"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">
                                                    Status
                                                </span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{type()} {status()}</span>
                                        </div>
                                    </li>


                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Users"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="truncate font-medium">Population</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{population()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Users"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="truncate font-medium">Teachers</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{ft_teacher()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Calendar"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="truncate font-medium">Enrollment</span>
                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{enrollment()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Image"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">NCES</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{ncesid()}</span>
                                        </div>
                                    </li>
                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Image"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">NAICS</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{naics_code()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Calendar"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="truncate font-medium">District ID</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{districtid()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Image"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">Shelter ID</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{shelter_id()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Shield"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="truncate font-medium">FIPS</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{countyfips()}</span>
                                        </div>
                                    </li>

                                    <li class="flex items-center justify-between py-4 pl-4 pr-5 text-sm/6">
                                        <div class="flex w-0 flex-1 items-center">
                                            <Icon name={"Image"} class="size-5 shrink-0 text-gray-400" />
                                            <div class="ml-4 flex min-w-0 flex-1 gap-2">
                                                <span class="capitalize truncate font-medium">{val_method().toLowerCase()}</span>

                                            </div>
                                        </div>
                                        <div class="ml-4 shrink-0">
                                            <span class="shrink-0 text-gray-400">{val_date()} </span>
                                        </div>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </dl>
                </div>
            </div>
            <div
                class="mx-auto grid max-w-md grid-cols-1 gap-6">
                <div class="space-y-6">
                    <section aria-labelledby="notes-title">
                        <div class="bg-gray-200 shadow ">
                            <div class="divide-y divide-gray-200">
                                <div class="px-4 py-5 sm:px-6">
                                    <h2 id="notes-title" class="text-lg font-medium text-gray-900">Notes</h2>
                                </div>
                                <div class="bg-white px-4 py-6 sm:px-6 text-left">
                                    <ul role="list" class="space-y-8">
                                        <li>
                                            <div class="flex space-x-3">
                                                <div class="shrink-0">
                                                    <img class="size-10 rounded-full"
                                                         src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                                                         alt=""/>
                                                </div>
                                                <div>
                                                    <div class="text-sm">
                                                        <a href="#" class="font-medium text-gray-900">Leslie
                                                            Alexander</a>
                                                    </div>
                                                    <div class="mt-1 text-sm text-gray-700">
                                                        <p>Ducimus quas delectus ad maxime totam doloribus
                                                            reiciendis ex. Tempore dolorem maiores. Similique
                                                            voluptatibus tempore non ut.</p>
                                                    </div>
                                                    <div class="mt-2 space-x-2 text-sm">
                                                        <span class="font-medium text-gray-500">4d ago</span>
                                                        <span class="font-medium text-gray-500">&middot;</span>
                                                        <button type="button"
                                                                class="font-medium text-gray-900">Reply
                                                        </button>
                                                    </div>
                                                </div>
                                            </div>
                                        </li>
                                        <li>
                                            <div class="flex space-x-3">
                                                <div class="shrink-0">
                                                    <img class="size-10 rounded-full"
                                                         src="https://images.unsplash.com/photo-1519244703995-f4e0f30006d5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                                                         alt=""/>
                                                </div>
                                                <div>
                                                    <div class="text-sm">
                                                        <a href="#" class="font-medium text-gray-900">Michael
                                                            Foster</a>
                                                    </div>
                                                    <div class="mt-1 text-sm text-gray-700">
                                                        <p>Et ut autem. Voluptatem eum dolores sint necessitatibus
                                                            quos. Quis eum qui dolorem accusantium voluptas
                                                            voluptatem ipsum. Quo facere iusto quia accusamus veniam
                                                            id explicabo et aut.</p>
                                                    </div>
                                                    <div class="mt-2 space-x-2 text-sm">
                                                        <span class="font-medium text-gray-500">4d ago</span>
                                                        <span class="font-medium text-gray-500">&middot;</span>
                                                        <button type="button"
                                                                class="font-medium text-gray-900">Reply
                                                        </button>
                                                    </div>
                                                </div>
                                            </div>
                                        </li>
                                        <li>
                                            <div class="flex space-x-3">
                                                <div class="shrink-0">
                                                    <img class="size-10 rounded-full"
                                                         src="https://images.unsplash.com/photo-1506794778202-cad84cf45f1d?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
                                                         alt=""/>
                                                </div>
                                                <div>
                                                    <div class="text-sm">
                                                        <a href="#" class="font-medium text-gray-900">Dries
                                                            Vincent</a>
                                                    </div>
                                                    <div class="mt-1 text-sm text-gray-700">
                                                        <p>Expedita consequatur sit ea voluptas quo ipsam
                                                            recusandae. Ab sint et voluptatem repudiandae voluptatem
                                                            et eveniet. Nihil quas consequatur autem. Perferendis
                                                            rerum et.</p>
                                                    </div>
                                                    <div class="mt-2 space-x-2 text-sm">
                                                        <span class="font-medium text-gray-500">4d ago</span>
                                                        <span class="font-medium text-gray-500">&middot;</span>
                                                        <button type="button"
                                                                class="font-medium text-gray-900">Reply
                                                        </button>
                                                    </div>
                                                </div>
                                            </div>
                                        </li>
                                    </ul>
                                </div>
                            </div>
                            <div class="bg-gray-50 px-4 py-6 sm:px-6">
                                <div class="flex space-x-3">
                                    <div class="shrink-0">
                                        <img class="size-10 rounded-full"
                                             src="https://images.unsplash.com/photo-1517365830460-955ce3ccd263?ixlib=rb-=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=8&w=256&h=256&q=80"
                                             alt=""/>
                                    </div>
                                    <div class="min-w-0 flex-1">
                                        <form action="#">
                                            <div>
                                                <label for="comment" class="sr-only">About</label>
                                                <textarea id="comment" name="comment" rows="3"
                                                          class="block w-full rounded-md border-0 px-3 py-1.5 shadow-sm outline-none ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-blue-600 sm:text-sm/6"
                                                          placeholder="Add a note"></textarea>
                                            </div>
                                            <div class="mt-3 flex items-center justify-between">
                                                <a href="#"
                                                   class="group inline-flex items-start space-x-2 text-sm text-gray-500 hover:text-gray-900">
                                                    <svg
                                                        class="size-5 shrink-0 text-gray-400 group-hover:text-gray-500"
                                                        viewBox="0 0 20 20" fill="currentColor" aria-hidden="true"
                                                        data-slot="icon">
                                                        <path fill-rule="evenodd"
                                                              d="M18 10a8 8 0 1 1-16 0 8 8 0 0 1 16 0ZM8.94 6.94a.75.75 0 1 1-1.061-1.061 3 3 0 1 1 2.871 5.026v.345a.75.75 0 0 1-1.5 0v-.5c0-.72.57-1.172 1.081-1.287A1.5 1.5 0 1 0 8.94 6.94ZM10 15a1 1 0 1 0 0-2 1 1 0 0 0 0 2Z"
                                                              clip-rule="evenodd"/>
                                                    </svg>
                                                    <span>Some HTML is okay.</span>
                                                </a>
                                                <button type="submit"
                                                        class="inline-flex items-center justify-center rounded-md bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600">Comment
                                                </button>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </section>
                </div>

            </div>

        </div>
    )
}

export {SchoolDetails}