from compose.cli.main import TopLevelCommand, project_from_options
import inspect
from lambdas import utils
from lambdas.getIdentifiersInfo import handler
from localstack.utils.aws import aws_stack
from localstack.utils.testutil import create_lambda_archive
import os
from dotenv import load_dotenv, find_dotenv

# Call & load env variables from .env file or environment
load_dotenv(find_dotenv(), override=True)


class TestInfoLambdaClass:
    options = {'--no-deps': False,
               '--abort-on-container-exit': False,
               '--scale': [], 'SERVICE': '',
               '--remove-orphans': False,
               '--no-recreate': True,
               '--force-recreate': False,
               '--build': False,
               '--no-build': False,
               '--no-color': False,
               '--rmi': None,
               '--volumes': '',
               '--follow': False,
               '--timestamps': False,
               '--tail': 'all',
               '-d': True
               }
    project = project_from_options(os.path.dirname(__file__), options)
    cmd = TopLevelCommand(project)

    def setUp(self):
        TestInfoLambdaClass.cmd.up(TestInfoLambdaClass.options)

    def tearDown(self):
        TestInfoLambdaClass.cmd.down(TestInfoLambdaClass.options)

    def test_invoke(self):
        lambda_client = aws_stack.connect_to_service('lambda')
        lambda_fp = inspect.getfile(handler)
        id_info_zip = utils.generate_lambda_zip(lambda_fp, write=True)
        lambda_client.create_function(FunctionName='f1',
                                      Runtime='python3.6',
                                      Role='r1',
                                      Handler='handler.handler',
                                      Code={'ZipFile': id_info_zip})
        response = lambda_client.invoke(FunctionName='f1')
        print(response['Payload'].read())
