az group create --name pocketnow --location westus3
az cosmosdb create --name pocketnow --resource-group pocketnow
az cosmosdb keys list --name pocketnow --resource-group pocketnow --query primaryMasterKey --output tsv
