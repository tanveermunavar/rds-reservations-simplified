# rdsReservationsSimplified

GoLang script to calculate the normalized units for a given instance class in the RDS in a specific. This helps to simply the reservation which otherwise is a manual and time consuming task. 

## Usage 

git clone https://github.com/tanveermunavar/rdsReservationsSimplified.git

export GOPATH=$(go env GOPATH)

go run calculate_rds_normalized_units.go -region ap-south-1 -instance db.m5


