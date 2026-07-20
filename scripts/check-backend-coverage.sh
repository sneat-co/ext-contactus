#!/usr/bin/env sh
# Keep executable Contactus extension helpers covered by focused,
# non-decreasing numeric floors.
set -eu

repo_root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
cd "$repo_root/backend"

check_coverage() {
	package=$1
	floor_file=$2
	label=$3
	floor=$(tr -d '[:space:]' < "$floor_file")
	coverage=$(go test -cover "$package" | awk '/coverage:/{ value=$(NF-2); sub(/%/, "", value); print value }')
	if [ -z "$coverage" ]; then
		echo "could not determine ${label} coverage" >&2
		exit 1
	fi
	if ! awk -v coverage="$coverage" -v floor="$floor" 'BEGIN { exit !(coverage >= floor) }'; then
		echo "${label} coverage ${coverage}% is below floor ${floor}%" >&2
		exit 1
	fi
	printf '%s coverage passed (coverage %s%%; floor %s%%).\n' "$label" "$coverage" "$floor"
}

check_coverage ./contactusmodels/briefs4contactus contactusmodels/briefs4contactus/coverage_floor.txt 'Contactus brief helpers'
check_coverage ./contactusmodels/const4contactus contactusmodels/const4contactus/coverage_floor.txt 'Contactus constants'
check_coverage ./dal4contactus dal4contactus/coverage_floor.txt 'Contactus record linkage helpers'
check_coverage ./facade4contactus facade4contactus/coverage_floor.txt 'Contactus facade linkage helpers'
