# Installation

## Build from source 
### Install golang 1.22 
installation for go in Linux, Windows , Macos in: https://go.dev/dl/

### Clonning the project from source
```sh
git clone https://github.com/aronyaina/workload-sim.git
cd workload-sim

go mod tidy
# Build bin app to be stress-test
# Syntax go build <binary_name> -o <go-main>
go build main -o main.go

# Run the stress test
main <file_name> <limit_request> <timeout_seconds> <requests_per_second>

```

## Run it from bin
```sh
main <file_name> <limit_request> <timeout_seconds> <requests_per_second>
```


# Context
C'est une application permettant de tester l'efficacite d'un serveur par l'envoie de requete, avec des limites definis par l'utilisateurs.
Golang est le language principal de cette application, pour des besoins multithread et ultra legere, facile a compile et tres portable 

# Utilisation
Nom du programme suivi des options suivant
```
main <file_name> <limit_request> <timeout_seconds> <requests_per_second>
```
### file_name 
determine le fichier ayant le format json , permettant de seeder les requetes:
```json
[
  {
    "url": "http://localhost:8080/projects",
    "method": "GET",
  },
  {
    "url": "http://localhost:8080/projects",
    "method": "POST",
    "body": {
      "project_name": "foo442901"
    }
  },
  {
    "url": "http://localhost:8080/projects/5",
    "method": "PUT",
    "form": {
      "project_name": "foo442901"
    }
  },
  {
    "url": "http://localhost:8080/projects/4",
    "method": "DELETE"
  },
]
```

### limit_request
le total des requete a executer en total, si le requete depasse la taille des requetes dans le fichier json, il sera alors rediriger a la longeur maximum de la requete
dans notre cas : limit_request = 3

### timeout_seconds
le temps accepter pour qu'une requete sois execute, si la duree d'une requete depasse ce temps, le programme s'arrete

### request_per_second
le total de requte par seconde pour une meilleur simulation

# Release
## V1.1 (Sans PUT - DELETE)
## V1.2 (Avec PUT - DELETE)
