podTemplate(
    label: 'apipoi',
    containers: [
        containerTemplate(name: 'az', image: 'microsoft/azure-cli:2.0.41', ttyEnabled: true, command: 'cat'),
    ]
) {
    node('apipoi'){
        stage('API-Poi-CI') {
            git 'https://github.com/Azure-Samples/openhack-devops-team.git'
            container('az'){
                stage('Docker Build POI API') {
                    withCredentials([azureServicePrincipal('azure_sp')]) {
                        sh """
                        az login --service-principal -u $AZURE_CLIENT_ID -p $AZURE_CLIENT_SECRET -t $AZURE_TENANT_ID
                        az acr login -n $acr_user -p $acr_pass -u $acr_user
                        az acr build --registry $acr_user -f Dockerfile --image devopsoh/api-hello:1.0  ./apis/poi/web
                        """
                    }
                }
            }
        }
    }
}
