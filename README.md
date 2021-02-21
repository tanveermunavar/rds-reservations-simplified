# RDS Reservations Simplified

GoLang script to calculate the normalized units for a given instance class in the RDS in a specific region. This helps to simply the reservation which is a manual and time consuming task. 

## Usage 

git clone https://github.com/tanveermunavar/rds-reservations-simplified.git

cd rds-reservations-simplified/

export GOPATH=$(go env GOPATH)

go run calculate_rds_normalized_units.go -region ap-south-1 -instance db.m5


