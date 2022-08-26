# customer_partner_match
This is a golang microservice for customers to find the right craftsman according to rating and distance.

An usecase for flooring partners experienced in wood, tiles and/or carpet has been implemented as a POC. Whatsoever, this microservice is prepare for new usecases implementation.


## Basic Info
- **Language:** Golang 1.18
- **Software Architecture:** Clean Architecture (ports and adaptors)
- **Infrastructure:** Postgres (with postGIS)
- Please find available for download a Postman request collection and the database.

## Requirements
- Golang [installation](https://www.youtube.com/watch?v=0dnTNElroug)
- Postgres with PostGIS (using PGAdmin 4) [installation](https://www.youtube.com/watch?v=LhKj-_-CCfY), [import database](https://www.youtube.com/watch?v=C30q5i1e9KE)
- Postman (optional) [installation](https://www.youtube.com/watch?v=3eHJkcA8mTs), [import collection](https://www.youtube.com/watch?v=bzquMXmCLUQ)

## Endpoints
### Register a new flooring craftsman
- path: "/floor-partner/new"
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
-----
