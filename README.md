
## THE MICROSERVICE
This customer_partner_match repository is a golang microservice for customers to find the right craftsman partner according to best rating and smallest distance.

A usecase for flooring partners experienced in wood, tiles and/or carpet has been implemented as a POC. Whatsoever, this microservice is prepare for new usecases implementation.


## Basic Info
- **Language:** Golang 1.18
- **Software Architecture:** Clean Architecture (ports and adapters)
- **Infrastructure:** Relational database (Postgres with postGIS)
- It's also available in this repository a Postman request collection (partners_floor.postman_collection.json) and a database (partners_location.sql).

## Requirements
- Golang [installation](https://www.youtube.com/watch?v=0dnTNElroug)
- Postgres with PostGIS (using PGAdmin 4) [installation](https://www.youtube.com/watch?v=LhKj-_-CCfY), [import database](https://www.youtube.com/watch?v=C30q5i1e9KE)
- Postman (optional) [installation](https://www.youtube.com/watch?v=3eHJkcA8mTs), [import collection](https://www.youtube.com/watch?v=bzquMXmCLUQ)

## Performance
When a client searches for partners the "best match" quality is determined first on
partner's average rating and second by customer-partner distance. Consequently, the distance between client and partner is calculated for each request since the client coordinates are previously unknown.

The solution with best performance for this microservice specification is the use of a spatial database in which the distance between client and partner is calculated using the ST_DistanceSphere function according to their coordinates. The database query also:
- considers the combination of flooring services the client wants (wood, tiles and/or carpet, i.e.),
- if the client location is in the partner's great circle operating area,
- best rating partners,
- max. total number of partners and
- the database returns the selected partners ordered by rating (desc) and client-partner distance (asc).

This results in a "lean" floor domain because multiple business rules are in the query. On the other hand, it would not make sense to migrate part of the business rules from the query to the domain and lose performance.

## Package structure
```
customer_partner_match/
|-- domain/
|     |-- floor/ 
|-- infrastrucutre/
|     |-- config/ 
|     |-- db/ 
|-- model/
|-- pkg/
|-- ports/
|     |-- input/
|     |-- output/
|-- web/
|     |-- v1/
|     |-- routes.go
|-- .env
|-- go.mod
|-- main.go
|-- README.md
```

# Endpoints
### Register a new flooring craftsman
- path: **"/floor-partner/new"**
- method: POST
- headers: none
- params: none
- body: 
  - **partner** (string): name of the partner/company,
  - **operating_radius** (int): partner's great circle ray (meters) regarding its location,
  - **latitude** (float64): geographical partner's coordinate (varying from -180 to 180),
  - **longitude** (float64): geographical partner's coordinate (varying from -180 to 180),
  - **wood** (bool): indicates whether this craftsman offers hardwood flooring (optional),
  - **tiles** (bool): indicates whether this craftsman offers tiles flooring (optional),
  - **carpet** (bool): indicates whether this craftsman offers carpet flooring (optional),


> After body fields validation, creates an id, transforms the partner's name to uppercase and attributes a default rating (4.0) to the craftsman.

Response

**status 200** (example):
```
{
    "data": [
        {
            "id": "d82ba7cf-527f-4cd5-8119-56ac9f8ca719",
            "partner": "KNOLL CRAFTSMAN",
            "rating": 4.6,
            "distance": 111,
            "latitude": -12.402,
            "longitude": 40.201
        },
        {
            "id": "6352bc86-6ed7-4c28-b34d-60dfd1385445",
            "partner": "WOODFLOORS COMPANY",
            "rating": 4.6,
            "distance": 2540,
            "latitude": -12.401298,
            "longitude": 40.202932
        },
	{
            "id": "4772bc86-6ed7-4c28-b34d-60dfd1381234",
            "partner": "TIMBER & TILES",
            "rating": 4.2,
            "distance": 245,
            "latitude": -12.402,
            "longitude": 40.2012
        }
    ]
}
OBS: distance in meters
```

**status 400**
- wrong input fields, invalid phone number (exemple):
```
{
    "err": "invalid phone number regex",
    "message": "missing/invalid field(s)",
    "errorKey": "error.InternalServerError",
    "status": 500
}
```
**status 500**
- database error
-----
### Search flooring partners
- path: **"/floor-partners"**
- method: GET
- headers: none
- params:
  - **latitude** (float): geographical customer's coordinate (varying from -180 to 180),
  - **longitude** (float): geographical customer's coordinate (varying from -180 to 180),
  - **floor_area** (float): total floor area (m²),
  - **phone** (string): customer's phone,
  - **total_partners** (int): indicates up to how many partners must be returned (optional, default value = 15),
  - **wood** (bool): returns craftsman with experience in hardwood flooring (optional),
  - **tiles** (bool): returns craftsman with experience in tiles flooring (optional),
  - **carpet** (bool): returns craftsman with experience in carpet flooring (optional),
- body: none

> After param fields validation, searches for best partners according to best rating and smallest distance. At least one flooring type (wood, tiles or carpet) needs to be true.

Response

**status 200** (example):
``` 
{
    "data": {
        "id": "153e46ed-8860-4afb-896e-a1ba490884d3",
        "partner": "TILES AND CARPET COMPANY",
        "rating": 4,
        "operating_radius": 30000,
        "latitude": 52.520008,
        "longitude": 13.404954,
        "wood": true,
        "carpet": true,
        "tiles": true
    }
}
```

**status 400**
- wrong input fields (exemple):
```
{
    "err": "latitude, longitude, floor_area",
    "message": "missing/invalid field(s)",
    "errorKey": "error.InputError",
    "status": 400
}
```
**status 500**
- database error
-----


# Installation
### Environment Variables (example)
```azure
DB_HOST=localhost
DB_PORT=5432
DB_NAME=partners_location
DB_USER=postgres
DB_PASSWORD=1234
API_PORT=85
```
### Running locally
A simple way to get started is to 
1. download this repository:
```
git clone https://github.com/yuripiffer/customer_partner_match.git
```
2. Start PGadmin 4 and import the database (or create your database),
3. Export the environment variables listed above (adapt if needed) and run the customer_partner-match microsservice.
5. Import the Postman collection and test the endpoints.

## Create your own database and tables

Create a spacial postGIS database (make sure to have postGIS installed):
```azure
CREATE {database name};
```

Add postGIS extention to the database:
```azure
CREATE EXTENSION postgis;
```

Create partner's tables:
```azure
CREATE TABLE {choose a table name} (
	id UUID primary key, 
	partner VARCHAR(255),
	rating numeric not null check (rating between 1 and 5), 
	operating_radius numeric CHECK (operating_radius > 0), 
	latitude numeric not null check (latitude between -180 and 180), 
	longitude numeric not null check (longitude between -180 and 180), 
	wood bool, 
	carpet bool, 
	tiles bool);
```

## License
This is an open-source and free software release.

