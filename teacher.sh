s=$(head -n 173 mystery/streets/Mattapan_Street |tail -n 1 | awk -F"#" '{print $2}')
echo $s
cat mystery/interviews/interview-$d
echo $MAIN_SUSPECT