# book-service 

## requirements
* go
* docker
* docker-compose
* gnu make
* protobuf-compiler

## startup 
1. `make run_db`
2. `make gen_proto`
3. `make build_and_run`

## testing 
`make test`

## Directories structure
* __cmd__ (app entrypoint _func main()_)
* __db__ (extended mysql dockerfile + init sql file)
* __internal__ (app source)
    * __models__ (sql models)
    * __repos__ (services who interact with sql)
    * __services__ (implementation grpc services from proto files)
    * __utils__ (app utilities)
* __proto__ (.proto files design and generated files)