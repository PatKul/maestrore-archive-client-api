#!bin/bash
# Create variable with blue and green color
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}Building Maestrore Archive${NC}"

# Get the start time in milliseconds
start=`date +%s`

# Check if bin folder exists if not create it
[ ! -d "./bin" ] &&
mkdir "./bin"

# Build the project
go build -o "./bin/maestrore-archive.exe"

end=`date +%s`

runtime=$((end-start))


# Check if the build was successful
if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build Successful${NC}"
    echo -e "${GREEN}Build Time: ${runtime} seconds${NC}"
else
    echo -e "${RED}Build Failed${NC}"
fi

