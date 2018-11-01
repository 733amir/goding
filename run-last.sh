#! /bin/sh

# Defining necessary variables
code_path="`find . -name \"*.go\" | sort | tail -1`"
code_dir="$(dirname ${code_path})"
bin_path="${code_dir}/binary"
code_input="${code_dir}/input"
code_expected="${code_dir}/expected"
code_output="${code_dir}/output"

go build -o "${bin_path}" "${code_path}"

time "${bin_path}" < "${code_input}" > "${code_output}"
printf "\n"

difference="`diff -iay --suppress-common-lines ${code_expected} ${code_output}`"
if [ -z "$difference" ]; then
    printf "Accepted\n"
else
    printf "${difference}\n"
fi