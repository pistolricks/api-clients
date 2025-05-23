# US Public Schools API

A CRUD API for US public schools data using PostgreSQL with a SolidJS client for displaying schools on a map using OpenLayers.

## Prerequisites

For the API:
- Go 1.24 or higher
- PostgreSQL 12 or higher with PostGIS extension
- Git

For the client:
- Node.js 14 or higher
- npm or yarn

## Setup

### API Setup

1. Clone the repository:
   ```
   git clone https://github.com/pistolricks/api-clients.git
   cd api-clients
   ```

2. Install Go dependencies:
   ```
   go mod download
   ```

3. Set up PostgreSQL:
   - Create a database named `schools`
   - Make sure the PostGIS extension is installed
   - Update the `.env` file with your database credentials if needed

4. Run the API:
   ```
   go run cmd/api/main.go
   ```
   or using the Makefile:
   ```
   make run/api
   ```

### Client Setup

1. Install client dependencies:
   ```
   make client/install
   ```
   or
   ```
   cd client
   npm install
   ```

2. Run the client:
   ```
   make client/dev
   ```
   or
   ```
   cd client
   npm run dev
   ```

   The client will be accessible at http://localhost:3003

3. To run both the API and client concurrently:
   ```
   make run/all
   ```

## Environment Variables

The application uses the following environment variables, which can be set in the `.env` file:

- `DB_HOST`: PostgreSQL host (default: localhost)
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_USER`: PostgreSQL username (default: postgres)
- `DB_PASSWORD`: PostgreSQL password (default: postgres)
- `DB_NAME`: PostgreSQL database name (default: schools)
- `DB_SSLMODE`: PostgreSQL SSL mode (default: disable)
- `PORT`: Server port (default: 8080)

## API Endpoints

### Schools

- `GET /api/schools`: List schools (with pagination)
  - Query parameters:
    - `page`: Page number (default: 1)
    - `pageSize`: Number of items per page (default: 10, max: 100)

- `GET /api/schools/{id}`: Get a school by ID

- `POST /api/schools`: Create a new school
  - Required fields:
    - `name`: School name
    - `latitude`: Latitude coordinate
    - `longitude`: Longitude coordinate

- `PUT /api/schools/{id}`: Update a school

- `DELETE /api/schools/{id}`: Delete a school

- `POST /api/schools/import`: Import schools from the GeoJSON file

## Data Model

The school data model includes the following fields:

- `id`: Unique identifier (auto-generated)
- `objectid`: Original object ID from the GeoJSON
- `name`: School name
- `address`: Street address
- `city`: City
- `state`: State
- `zip`: ZIP code
- `country`: Country
- `county`: County
- `countyfips`: County FIPS code
- `latitude`: Latitude coordinate
- `longitude`: Longitude coordinate
- `level`: School level (e.g., ELEMENTARY, MIDDLE, HIGH)
- `st_grade`: Starting grade
- `end_grade`: Ending grade
- `enrollment`: Number of students
- `ft_teacher`: Number of full-time teachers
- `type`: School type
- `status`: School status
- `population`: Population
- `ncesid`: NCES ID
- `districtid`: District ID
- `naics_code`: NAICS code
- `naics_desc`: NAICS description
- `website`: School website
- `telephone`: School telephone
- `sourcedate`: Source date
- `val_date`: Validation date
- `val_method`: Validation method
- `source`: Data source
- `shelter_id`: Shelter ID
- `created_at`: Record creation timestamp
- `updated_at`: Record update timestamp

## Example Requests

### List Schools

```
GET /api/schools?page=1&pageSize=10
```

### Get School by ID

```
GET /api/schools/1
```

### Create School

```json
POST /api/schools
{
  "name": "New School",
  "address": "123 Main St",
  "city": "Anytown",
  "state": "CA",
  "zip": "12345",
  "latitude": 37.7749,
  "longitude": -122.4194,
  "level": "ELEMENTARY",
  "st_grade": "K",
  "end_grade": "5"
}
```

### Update School

```json
PUT /api/schools/1
{
  "name": "Updated School Name",
  "address": "456 Oak St"
}
```

### Import Schools from GeoJSON

```
POST /api/schools/import
```

This will import all schools from the `us-public-schools.geojson` file.

## Client Application

The project includes a SolidJS client application that displays schools on a map using OpenLayers. The client runs on port 3003 and can be accessed at http://localhost:3003 when started.

### Client Features

- Interactive map displaying school locations
- Click on a school marker to view details
- Pagination to navigate through all schools
- Responsive design for desktop and mobile devices

### Building for Production

To build the client for production using the Makefile:

```
make client/build
```

Or using npm directly:

```
cd client
npm run build
```

The built files will be in the `client/dist` directory and can be served by any static file server.
