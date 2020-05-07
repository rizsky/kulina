# Basic Concepts #
1. Codesmell
example smell : non-seperation concern of layers, code scatter everywhere and all logic (business logic, delivery logic, and etc) just group in one place or two, this is cause of the software hard to unit test or maintenance because high coupling for the functions itself

how to fix this : create layer independents, easy to maintenance and refactorable  

- controller/handler layer
- service layer
- repository layer
- model layer

2. Dep Injections
in golang :
dependancy injections the components should receive their dependancies when being created. futhermore in golang we can take the sample below if the config is interface we can mock for testing.

sample :
type Database struct {
  config *Config
}

func NewDB(config *Config) *Database {
  return &Database{
    config: config,
  }
}

the purpose is Dependency Injection helps build more robust and testable applications



# Rest API #

based on standard of REST below the common using the method

POST
- do : create a new resource, 
- donts : doing non-manipulating resource or delete & update resource, because it will give confusing to whom consuming the api

GET
- do : fetch data, using query string to filtering the data
- donts : manipulating resource, because when doing this its gonna be a confusing for whom consuming the api