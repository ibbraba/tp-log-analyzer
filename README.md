# TP : GoLog Analyzer - Analyse de Logs Distribuée



## Build du projet 

A la racine du projet exécutez la commande suivante : 

``` go build -o bin\log-analyzer.exe ``` 

OU 

``` go build -o bin\log-analyzer ```  Sur MacOS


## Lancer Le script 

Lancez le script avec la commande suivante : 

``` bin\log-analyzer.exe analyze -p config.json ``` 

Avec le flag -p indiquant le nom de votre ficher à lire 
Par exemple si votre fichier est nommé log.json exécutez 
``` bin\log-analyzer.exe analyze -p log.json ``` 

## Modifier le fichier de destination 

Par defaut le script crée un fichier de report nommé report.json 

Pour changer le nom ajoutez le flag -o par exemple 

``` bin\log-analyzer.exe analyze -p config.json -o custom_output.json ```



## Bonus : Filtrage par statut

Ajoutez le flag -s pour filter les logs par statut (OK ou FAILED) par exemple 

``` bin\log-analyzer.exe analyze -p config.json -s ok -o oklogs.json ``` 
pour obtenir les logs ayant reussi l'analyse

OU

``` bin\log-analyzer.exe analyze -p config.json -s failed -o failedlogs.json  ```
pour obtenir les logs ayant échoué

Enjoy !