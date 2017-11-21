#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# This script pins the BCCSP package family from Hyperledger Fabric into the SDK
# These files are checked into internal paths.
# Note: This script must be adjusted as upstream makes adjustments

IMPORT_SUBSTS=($IMPORT_SUBSTS)

GOIMPORTS_CMD=goimports
GOFILTER_CMD="go run scripts/_go/cmd/gofilter/gofilter.go"

declare -a PKGS=(
    "api"
    "lib"
    "lib/tls"
    "sdkpatch/logbridge"
    "util"
)

declare -a FILES=(
    "api/client.go"
    "api/net.go"

    "lib/client.go"
    "lib/identity.go"
    "lib/signer.go"
    "lib/clientconfig.go"
    "lib/util.go"
    "lib/sdkpatch_serverstruct.go"

    "lib/tls/tls.go"

    "sdkpatch/logbridge/logbridge.go"
    "sdkpatch/logbridge/syslogwriter.go"

    "util/util.go"
    "util/csp.go"
)

echo 'Removing current upstream project from working directory ...'
rm -Rf "${INTERNAL_PATH}"
mkdir -p "${INTERNAL_PATH}"

# Create directory structure for packages
for i in "${PKGS[@]}"
do
    mkdir -p $INTERNAL_PATH/${i}
done

# Apply fine-grained patching
gofilter() {
    echo "Filtering: ${FILTER_FILENAME}"
    cp ${TMP_PROJECT_PATH}/${FILTER_FILENAME} ${TMP_PROJECT_PATH}/${FILTER_FILENAME}.bak
    $GOFILTER_CMD -filename "${TMP_PROJECT_PATH}/${FILTER_FILENAME}.bak" \
        -filters "$FILTERS_ENABLED" -fn "$FILTER_FN" -gen "$FILTER_GEN" -mode "$FILTER_MODE" \
        > "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
} 

echo "Filtering Go sources for allowed functions ..."
FILTER_MODE="allow"
FILTERS_ENABLED="fn"

FILTER_FILENAME="lib/client.go"
FILTER_FN="Enroll,GenCSR,SendReq,Init,newPost,newEnrollmentResponse,newCertificateRequest"
FILTER_FN+=",getURL,NormalizeURL,initHTTPClient,net2LocalServerInfo,NewIdentity"
gofilter
sed -i'' -e 's/util.GetServerPort()/\"\"/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/log "github.com\// a\
"github.com\/hyperledger\/fabric-sdk-go\/api\/apicryptosuite"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.BCCSP/apicryptosuite.CryptoSuite/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.Key/apicryptosuite.Key/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"

FILTER_FILENAME="lib/identity.go"
FILTER_FN="newIdentity,Revoke,Post,addTokenAuthHdr,GetECert,Reenroll,Register,GetName"
gofilter
sed -i'' -e 's/util.GetDefaultBCCSP()/nil/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/log "github.com\// a\
"github.com\/hyperledger\/fabric-sdk-go\/api\/apicryptosuite"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.BCCSP/apicryptosuite.CryptoSuite/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.Key/apicryptosuite.Key/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"

FILTER_FILENAME="lib/signer.go"
FILTER_FN="newSigner,Key,Cert"
gofilter
sed -i'' -e '/"github.com\// a\
"github.com\/hyperledger\/fabric-sdk-go\/api\/apicryptosuite"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.Key/apicryptosuite.Key/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"

FILTER_FILENAME="lib/clientconfig.go"
FILTER_FN=
gofilter

FILTER_FILENAME="lib/util.go"
FILTER_FN="GetCertID,BytesToX509Cert"
gofilter

FILTER_FILENAME="lib/tls/tls.go"
FILTER_FN="GetClientTLSConfig,AbsTLSClient,checkCertDates"
gofilter
sed -i'' -e '/log "github.com\// a\
"github.com\/hyperledger\/fabric-sdk-go\/api\/apicryptosuite"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/"crypto\// a\
cryptosuite "github.com\/hyperledger\/fabric-sdk-go\/pkg\/cryptosuite\/bccsp"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.BCCSP/apicryptosuite.CryptoSuite/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/csp = factory.GetDefault()/csp = cryptosuite.GetSuite(factory.GetDefault())/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"


FILTER_FILENAME="util/csp.go"
FILTER_FN="InitBCCSP,ConfigureBCCSP,GetBCCSP,makeFileNamesAbsolute"
FILTER_FN+=",getBCCSPKeyOpts,ImportBCCSPKeyFromPEM,LoadX509KeyPair,GetSignerFromCert"
FILTER_FN+=",BCCSPKeyRequestGenerate,GetSignerFromCertFile"
gofilter
sed -i'' -e '/_.\"time\"/d' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/\"github.com\/cloudflare\/cfssl\/cli\"/d' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/\"github.com\/cloudflare\/cfssl\/ocsp\"/d' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/log "github.com\// a\
"github.com\/hyperledger\/fabric-sdk-go\/api\/apicryptosuite"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e '/cspsigner "github.com\// a\
cryptosuite "github.com\/hyperledger\/fabric-sdk-go\/pkg\/cryptosuite\/bccsp"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.BCCSP/apicryptosuite.CryptoSuite/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.Key/apicryptosuite.Key/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
# skip first substitution of "return csp, nil"
START_LINE=`grep -n "return csp, nil" "${TMP_PROJECT_PATH}/${FILTER_FILENAME}" | head -n 1 | awk -F':' '{print $1}'`
START_LINE=$((START_LINE+1))
sed -i'' -e ${START_LINE}',$ s/return csp, nil/return cryptosuite.GetSuite(csp), nil/g'  "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"

FILTER_FILENAME="util/util.go"
FILTER_FN="ReadFile,HTTPRequestToString,HTTPResponseToString"
FILTER_FN+=",GetX509CertificateFromPEM,GetSerialAsHex,GetEnrollmentIDFromPEM"
FILTER_FN+=",MakeFileAbs,Marshal,StructToString,LoadX509KeyPair,CreateToken"
FILTER_FN+=",GenECDSAToken,GetEnrollmentIDFromX509Certificate,B64Encode,B64Decode"
FILTER_FN+=",GetMaskedURL"
gofilter
sed -i'' -e '/log "golang.org\/x/ a\
"github.com\/hyperledger\/fabric-sdk-go\/api\/apicryptosuite"\
' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.BCCSP/apicryptosuite.CryptoSuite/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"
sed -i'' -e 's/bccsp.Key/apicryptosuite.Key/g' "${TMP_PROJECT_PATH}/${FILTER_FILENAME}"

# Apply patching
echo "Patching import paths on upstream project ..."
WORKING_DIR=$TMP_PROJECT_PATH FILES="${FILES[@]}" IMPORT_SUBSTS="${IMPORT_SUBSTS[@]}" scripts/third_party_pins/common/apply_import_patching.sh

echo "Inserting modification notice ..."
WORKING_DIR=$TMP_PROJECT_PATH FILES="${FILES[@]}" scripts/third_party_pins/common/apply_header_notice.sh


# Copy patched project into internal paths
echo "Copying patched upstream project into working directory ..."
for i in "${FILES[@]}"
do
    TARGET_PATH=`dirname $INTERNAL_PATH/${i}`
    cp $TMP_PROJECT_PATH/${i} $TARGET_PATH
done
