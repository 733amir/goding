#! /bin/sh

# Defining necessary variables
code_path="`find . -name \"*.go\" -mindepth 2 | sort | tail -1`"
#echo "code_path = ${code_path}"
code_dir="$(dirname ${code_path})"
#echo "code_dir = ${code_dir}"
bin_path="${code_dir}/binary"
#echo "bin_path = ${bin_path}"
code_input="${code_dir}/input-"
#echo "code_input = ${code_input}"
code_expected="${code_dir}/expected-"
#echo "code_expected = ${code_expected}"
code_output="${code_dir}/output-"
#echo "code_output = ${code_output}"

go build -o "${bin_path}" "${code_path}"

i=0
for input in `ls -1 ${code_input}*`; do
    ${bin_path} < ${input} > ${code_output}${i}
    ((i++))
done

i=0
for output in `ls -1 ${code_output}*`; do
    difference="`diff -iay --suppress-common-lines ${code_expected}${i} ${output}`"
    if [[ ! -z "$difference" ]]; then
        printf "${difference}\n"
        exit 1
    fi
    ((i++))
done

printf "Accepted\n"
