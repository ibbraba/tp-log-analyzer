# TP : GoLog Analyzer - Analyse de Logs Distribuée



## Build du projet 

A la racine du projet exécutez la commande suivante : 

``` go build -o log-analyzer.exe ``` 

OU 

``` go build -o log-analyzer ```  Sur MacOS


## Lancer Le script 

Lancez le script avec la commande suivante : 

``` .\log-analyzer.exe analyze -p config.json ``` 

Avec le flag -p indiquant le nom de votre ficher à lire 
Par exemple si votre fichier est nommé log.json exécutez 
``` .\log-analyzer.exe analyze -p log.json ``` 

## Modifier le fichier de destination 

Par defaut le script crée un fichier de report nommé report.json 

Pour changer le nom ajoutez le flag -o par exemple 

``` .\log-analyzer.exe analyze -p config.json -o custom_output.json ```

Enjoy !