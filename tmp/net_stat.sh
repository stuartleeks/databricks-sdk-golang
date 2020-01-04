echo "ESTABLISHED: $(netstat -t | grep localhost | grep EST | wc -l)"
echo "TIME_WAIT: $(netstat -t | grep localhost | grep TIME_WAIT | wc -l)"
echo "TOTAL: $(netstat -t | grep localhost | wc -l)"
echo "--------------"
netstat -t | grep localhost