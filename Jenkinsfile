podTemplate(
    label: 'apipoi',
    containers: [
        containerTemplate(name: 'az', image: 'microsoft/azure-cli:2.0.41', ttyEnabled: true, command: 'cat'),
    ]
) {
    node('apipoi'){
        stage('API-Poi-CI') {
            git 'https://github.com/izpavlovich/Team4Power.git'
            container('az'){
                stage('Docker Build POI API') {
                    withCredentials([azureServicePrincipal('azure_sp')]) {
                        sh """
                        az login --service-principal -u $AZURE_CLIENT_ID -p $AZURE_CLIENT_SECRET -t $AZURE_TENANT_ID
                        az acr build --registry ohdrteam04acr2 -f Dockerfile --image devopsoh/api-hello:1.0  ./apis/poi/web
                        """
                    }
                }
            }
        }

        stage ('compile and test') {

            container('az') {
                sh "$WORKSPACE/apis/poi/build_deploy_poi.sh"
            }


        }
    }
}
