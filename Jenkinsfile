podTemplate(
    label: 'apipoi',
    containers: [
        containerTemplate(name: 'az', image: 'microsoft/azure-cli:2.0.41', ttyEnabled: true, command: 'cat'),
        containerTemplate(name: 'helm-builder', image: 'dtzar/helm-kubectl:2.9.1', ttyEnabled: true, command: 'cat'),

    ],
    envVars: [
            envVar(key: 'ACR_NAME', value: 'ohdrteam04acr2'),
            envVar(key: 'INGRESS_URL', value: 'akstraefikohdrteam04.westus.cloudapp.azure.com'),
            envVar(key: 'AKS_RG', value: 'ohdrteam04rg'),
            envVar(key: 'AKS_NAME', value: 'ohdrteam04aks'),
            envVar(key: 'IMAGE_NAME', value: 'devopsoh/api-poi-hellow')
    ]

) {
    node('apipoi'){
        stage('API-Poi CI') {
            git 'https://github.com/izpavlovich/Team4Power.git'
            container('az'){
                stage('Docker Build POI API') {
                    withCredentials([azureServicePrincipal('azure_sp')]) {
                        sh """
                        az login --service-principal -u $AZURE_CLIENT_ID -p $AZURE_CLIENT_SECRET -t $AZURE_TENANT_ID
                        az aks get-credentials -g $AKS_RG -n $AKS_NAME
                        az acr build --registry $ACR_NAME -f Dockerfile --image $IMAGE_NAME:${env.BUILD_ID}  ./apis/poi/web
                        """
                    }
                }
            }
        }

        stage ('Compile and test') {
            container('helm-builder'){
                sh """
                cd $WORKSPACE
                helm init
                helm install --name api-poi-hello --set repository.image=${ACR_NAME}.azurecr.io/${IMAGE_NAME},repository.tag=${env.BUILD_ID},ingress.rules.endpoint.host=$INGRESS_URL
                """
            }
        }
    }
}
