from dotenv import load_dotenv, find_dotenv
import json
import logging
import os
import requests
from requests.packages.urllib3.util.retry import Retry

# Call & load env variables from .env file or environment
load_dotenv(find_dotenv(), override=True)
API_ENV = os.environ['API_ENV']
DRUID_USERNAME = os.environ['DRUID_USERNAME']
DRUID_PWORD = os.environ['DRUID_PASSWORD']

# Start logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)

# Set global SURI HTTP call parameters
suri_url = 'https://{0}:{1}@sul-lyberservices-{2}.stanford.edu/suri2/namespaces/druid/identifiers?quantity={3}&response=text'
retry = Retry(
    # Maximum number of connection retries
    total=5,
    # backoff factor is in milliseconds
    backoff_factor=.01
    )
retries = requests.adapters.HTTPAdapter(max_retries=retry)


def handler(event, context):
    """Mint & return given number of DRUIDs, with 1 returned as default."""
    identifiers = []

    # Check for quantity of DRUIDs wanted, otherwise, default to 1
    if event.get("queryStringParameters").get("quantity"):
        quantity = int(event["queryStringParameters"]["quantity"])
    else:
        quantity = 1

    # Mint DRUID(s)
    druid_url = suri_url.format(DRUID_USERNAME, DRUID_PWORD, API_ENV, quantity)
    druid_url_safe = suri_url.format(DRUID_USERNAME, 'password', API_ENV, quantity)
    logger.info('Starting DRUID mint call: {0}'.format(druid_url_safe))

    try:
        session = requests.Session()
        session.mount('https://', retries)
        resp = session.post(druid_url).content
        logger.info('Parsing SURI response for new DRUID(s).')
        print(resp)
        return({'body': resp})
    except requests.exceptions.RequestException as e:
        exception_class = e.__class__.__name__
        exception_msg = e
        handle_network_errors(exception_class, exception_msg)


def handle_network_errors(exception_class, exception_msg):
    api_exception_obj = {
        'isError': True,
        'type': exception_class,
        'message': exception_msg
    }

    api_exception_json = json.dumps(api_exception_obj)
    raise InfoLambdaException(api_exception_json)


class InfoLambdaException(Exception):
    pass
