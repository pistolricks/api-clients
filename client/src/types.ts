// Define types for the school data
export interface School {
  id: number;
  objectid: number;
  name: string;
  address?: string;
  city?: string;
  state?: string;
  zip?: string;
  country?: string;
  county?: string;
  countyfips?: string;
  latitude: number;
  longitude: number;
  level?: string;
  st_grade?: string;
  end_grade?: string;
  enrollment?: number;
  ft_teacher?: number;
  type?: number;
  status?: number;
  population?: number;
  ncesid?: string;
  districtid?: string;
  naics_code?: string;
  naics_desc?: string;
  website?: string;
  telephone?: string;
  sourcedate?: string;
  val_date?: string;
  val_method?: string;
  source?: string;
  shelter_id?: string;
  created_at: string;
  updated_at: string;
}

// Define types for the API response
export interface SchoolsResponse {
  schools: School[];
  total: number;
  page: number;
  pageSize: number;
}

// Define types for map elements
export interface MapElements {
  mapElement: HTMLDivElement | undefined;
  popupElement: HTMLDivElement | undefined;
}