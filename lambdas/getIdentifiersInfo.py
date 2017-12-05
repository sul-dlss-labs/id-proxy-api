from dotenv import load_dotenv, find_dotenv
import json
import logging
import os
import requests
from requests.packages.urllib3.util.retry import Retry
from xml.etree import ElementTree

# Call & load env variables from .env file or environment
load_dotenv(find_dotenv(), override=True)
API_ENV = os.environ['API_ENV']
DRUID_USERNAME = os.environ['DRUID_USERNAME']
DRUID_PWORD = os.environ['DRUID_PASSWORD']

# Start logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)

# Set global SURI HTTP call parameters
suri_url = 'https://{0}:{1}@sul-lyberservices-{2}.stanford.edu/suri2/'
suri_ns_url = suri_url + 'namespaces/druid'
retry = Retry(
    # Maximum number of connection retries
    total=5,
    # backoff factor is in milliseconds
    backoff_factor=.01
    )
retries = requests.adapters.HTTPAdapter(max_retries=retry)


def handler(event, context):
    """Retrieve namespace / source information from ID systems proxied."""
    sources = []

    # Handle DRUIDs
    druid_ns_url = suri_ns_url.format(DRUID_USERNAME, DRUID_PWORD, API_ENV)
    druid_ns_url_safe = suri_ns_url.format(DRUID_USERNAME, 'password', API_ENV)
    logger.info('Starting DRUID Namespace call: {0}'.format(druid_ns_url_safe))

    try:
        session = requests.Session()
        session.mount('https://', retries)
        resp = session.get(druid_ns_url)
        tree = ElementTree.fromstring(resp.content)
        logger.info('Parsing SURI response for DRUID Namespace.')
        for ns in tree.iter('namespace'):
            source = {}
            source['name'] = ns.find('name').text
            source['template'] = ns.find('template').text
            sources.append(source)
        return({'body': sources})
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
