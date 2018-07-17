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
                        az login --service-principal -u 7aca892e-f9bd-4eba-84c2-636e7273a526 -p 9543f660db7ee1b925b3 --tenant f936f040-917b-4d70-a0c5-277ca7322c9e
                        az acr build --registry ohdrteam04acr -f Dockerfile --image devopsoh/poi-hello  ./apis/poi/web
                        """
                    }
                }
            }
        }
    }
}
