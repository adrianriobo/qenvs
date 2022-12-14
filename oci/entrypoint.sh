#!/bin/sh

VALID_CONFIG=true
# Check required ENVs
if [ -z "${OPERATION+x}" ]; then 
  echo "OPERATION is required"
  VALID_CONFIG=false
fi

if [ -z "${PROJECT_NAME+x}" ]; then 
  echo "PROJECT_NAME ENV is required"
  VALID_CONFIG=false  
fi

if [ -z "${BACKED_URL+x}" ]; then 
  echo "${INTERNAL_OUTPUT} will be used as backed url it will be exported as volume"
  BACKED_URL="file://${INTERNAL_OUTPUT}"
fi

if [ -z "${CONNECTION_OUTPUT+x}" ]; then 
  echo "${INTERNAL_OUTPUT} will be used as output folder for connecion resources"
  CONNECTION_OUTPUT="${INTERNAL_OUTPUT}"
fi

if [ -z "${AWS_ACCESS_KEY_ID+x}" ] || [ -z "${AWS_SECRET_ACCESS_KEY+x}" ] || [ -z "${AWS_DEFAULT_REGION+x}" ]; then 
  echo "AWS ENV for credentials are required"
  VALID_CONFIG=false  
fi

if [ -z "${PULUMI_CONFIG_PASSPHRASE+x}" ]; then 
  # https://www.pulumi.com/docs/reference/cli/environment-variables/
  PULUMI_CONFIG_PASSPHRASE="passphrase"
fi

if [ "${VALID_CONFIG}" = false ]; then
  echo "Add the required ENVs"
  exit 1
fi

# //https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
AWS_SDK_LOAD_CONFIG=1

if [[ "${OPERATION}" == "create" ]]; then
  if [ -z "${SUPPORTED_HOST_ID+x}" ]; then 
    echo "SUPPORTED_HOST_ID is required"
    VALID_CONFIG=false
  fi
  AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
  AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
  AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} \
  AWS_SDK_LOAD_CONFIG=${AWS_SDK_LOAD_CONFIG} \
  PULUMI_CONFIG_PASSPHRASE=${PULUMI_CONFIG_PASSPHRASE} \
        qenvs host create \
          --project-name "${PROJECT_NAME}" \
          --backed-url "${BACKED_URL}" \
          --conn-details-output "${CONNECTION_OUTPUT}" \
          --host-id "${SUPPORTED_HOST_ID}" \
          --rh-subscription-username "${RH_SUBSCRIPTION_USERNAME}" \
          --rh-subscription-password "${RH_SUBSCRIPTION_PASSWORD}"
else 
  AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
  AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
  AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} \
  AWS_SDK_LOAD_CONFIG=${AWS_SDK_LOAD_CONFIG} \
  PULUMI_CONFIG_PASSPHRASE=${PULUMI_CONFIG_PASSPHRASE} \
        qenvs host destroy \
          --project-name "${PROJECT_NAME}" \
          --backed-url "${BACKED_URL}" 
fi
