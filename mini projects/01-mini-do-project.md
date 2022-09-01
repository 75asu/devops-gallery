# A static HTML web app deployed in Azure App Service


[Live Demo](http://expanding-cards.azurewebsites.net)<br>
Here is the project which I have deployed using Azure App services - [Expanding Cards](https://github.com/measutosh/frontend-gallery/tree/main/expanding-cards)

Steps followed :- 
  - log into Azure CLI
  - clone the repo to a directory
  - enter into the repo and run the command `az webapp up --location westeurope --name <app_name> --html`
  - the result should like this
  ```
  asutosh [ ~/devopsprojects ]$ mkdir static-site-azure-app-services
  asutosh [ ~/devopsprojects ]$ cd static-site-azure-app-services/
  asutosh [ ~/devopsprojects/static-site-azure-app-services ]$ git clone https://gitlab.com/repocentral/50days50projects/12_expanding_cards.git
  Cloning into '12_expanding_cards'...
  remote: Enumerating objects: 8, done.
  remote: Counting objects: 100% (8/8), done.
  remote: Compressing objects: 100% (7/7), done.
  remote: Total 8 (delta 0), reused 0 (delta 0), pack-reused 0
  Receiving objects: 100% (8/8), 4.18 KiB | 1.39 MiB/s, done.
  asutosh [ ~/devopsprojects/static-site-azure-app-services ]$ cd 12_expanding_cards/
  asutosh [ ~/devopsprojects/static-site-azure-app-services/12_expanding_cards ]$ az webapp up --location westeurope --name expanding-cards --html
  The webapp 'expanding-cards' doesn't exist
  Creating Resource group '1844439_rg_4691' ...
  Resource group creation complete
  Creating AppServicePlan '1844439_asp_8347' ...
  Creating webapp 'expanding-cards' ...
  Configuring default logging for the app, if not already enabled
  Creating zip with contents of dir /home/asutosh/devopsprojects/static-site-azure-app-services/12_expanding_cards ...
  Getting scm site credentials for zip deployment
  Starting zip deployment. This operation can take a while to complete ...
  Deployment endpoint responded with status code 202
  You can launch the app at http://expanding-cards.azurewebsites.net
  Setting 'az webapp up' default arguments for current directory. Manage defaults with 'az configure --scope local'
  --resource-group/-g default: 1844439_rg_4691
  --sku default: F1
  --plan/-p default: 1844439_asp_8347
  --location/-l default: westeurope
  --name/-n default: expanding-cards
  {
    "URL": "http://expanding-cards.azurewebsites.net",
    "appserviceplan": "1844439_asp_8347",
    "location": "westeurope",
    "name": "expanding-cards",
    "os": "Windows",
    "resourcegroup": "1844439_rg_4691",
    "runtime_version": "-",
    "runtime_version_detected": "-",
    "sku": "FREE",
    "src_path": "//home//asutosh//devopsprojects//static-site-azure-app-services//12_expanding_cards"
  }
  ```
