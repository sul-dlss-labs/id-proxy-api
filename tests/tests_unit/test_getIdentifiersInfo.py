from localstack.utils.aws import aws_stack
from compose.cli.main import TopLevelCommand, project_from_options
import shutil
import os

id_lambda = '/Users/sul.cmharlow/Code/src/github.org/sul-dlss-labs/id-proxy-api/lambdas/getIdentifiersInfo.zip'
requirements = '/Users/sul.cmharlow/Code/src/github.org/sul-dlss-labs/id-proxy-api/requirements.txt'

class TestLambdaClass:
    options = {"--no-deps": False,
           "--abort-on-container-exit": False,
           "--scale": [],
           "SERVICE": "",
           "--remove-orphans": False,
           "--no-recreate": True,
           "--force-recreate": False,
           "--build": False,
           "--no-build": False,
           "--no-color": False,
           "--rmi": "none",
           "--volumes": "",
           "--follow": False,
           "--timestamps": False,
           "--tail": "all",
           "-d": True}
    project = project_from_options(os.path.dirname(__file__), options)
    cmd = TopLevelCommand(project)

    def setUp(self):
        TestLambdaClass.cmd.up(TestLambdaClass.options)

    def tearDown(self):
        TestLambdaClass.cmd.down(TestLambdaClass.options)

    def test_invoke(self):
        lambda_client = aws_stack.connect_to_service('lambda')
        lambda_client.create_function(
        FunctionName='f1',
        Runtime='python3.6',
        Role='r1',
        Handler='getIdentifiersInfo.handler',
        Code={'ZipFile': open(id_lambda, 'rb').read()}
        )
        response = lambda_client.invoke(FunctionName='f1')
        print(response['Payload'].read())
