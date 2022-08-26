[UNDER CONSTRUCTION]
## customer_partner_match microservice
This is a golang microservice for customers to find the right craftsman according to rating and distance.

An usecase for flooring partners experienced in wood, tiles and/or carpet has been implemented as a POC. Whatsoever, this microservice is prepare for new usecases implementation.


## Basic Info
- **Language:** Golang 1.18
- **Software Architecture:** Clean Architecture (ports and adaptors)
- **Infrastructure:** Relational database (Postgres with postGIS)
- Please find available for download a Postman request collection and the database.

## Requirements
- Golang [installation](https://www.youtube.com/watch?v=0dnTNElroug)
- Postgres with PostGIS (using PGAdmin 4) [installation](https://www.youtube.com/watch?v=LhKj-_-CCfY), [import database](https://www.youtube.com/watch?v=C30q5i1e9KE)
- Postman (optional) [installation](https://www.youtube.com/watch?v=3eHJkcA8mTs), [import collection](https://www.youtube.com/watch?v=bzquMXmCLUQ)

## Performance
abc Why a spacial database and why everything in the query...
sphera

## Package structure


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


> abc abc abc

*Response*

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

> abc abc abc

*Response*

-----


# Installation
### Environment Variables
```azure
DB_HOST=localhost
DB_PORT=5432
DB_NAME=partners
DB_USER=postgres
DB_PASSWORD=1234
```
### Running locally
A simple way to get started is to 
1. download this repository:
```
git clone https://github.com/yuripiffer/customer_partner-match.git
```
... database....
2. Export the environment variables listed above.
3. Run the customer_partner-match.
4. 

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

