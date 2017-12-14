import inspect
from localstack.utils.testutil import create_lambda_archive
from localstack.constants import LOCALSTACK_ROOT_FOLDER
import os
import pdb


def generate_libs():
    libs = []
    root_folder = LOCALSTACK_ROOT_FOLDER
    for library in os.listdir(root_folder):
        if 'dist-info' not in library:
            libs.append(library)
    return(libs)


def generate_envs():
    """Work around due to dotenv not working on lambda docker."""
    envs = {}
    local_envs = os.environ
    envs['API_ENV'] = local_envs['API_ENV']
    envs['DRUID_USERNAME'] = local_envs['DRUID_USERNAME']
    envs['DRUID_PWORD'] = local_envs['DRUID_PASSWORD']
    return(envs)


def generate_lambda_zip(lambda_filepath, write=True):
    # NB: localstack create_lambda_archive renames file handler.py
    libs = generate_libs()
    lambda_zip = create_lambda_archive(open(lambda_filepath).read(), libs=libs,
                                       runtime='python3.6')
    if write:
        zip_filepath = lambda_filepath.replace('.py', '.zip')
        with open(zip_filepath, 'wb') as fout:
            fout.write(lambda_zip)
    return(lambda_zip)
