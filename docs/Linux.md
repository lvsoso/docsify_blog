```shell
watch -n 1 "cat /proc/vmstat | grep dirty"

perf record -a -g -p <pid>

perf report

cd /sys/kernel/debug/tracing
echo vfs_write >> set_ftrace_filter
echo xfs_file_write_iter >> set_ftrace_filter
echo xfs_file_buffered_aio_write >> set_ftrace_filter
echo iomap_file_buffered_write
echo iomap_file_buffered_write >> set_ftrace_filter
echo pagecache_get_page >> set_ftrace_filter
echo try_to_free_mem_cgroup_pages >> set_ftrace_filter
echo try_charge >> set_ftrace_filter
echo mem_cgroup_try_charge >> set_ftrace_filter
echo function_graph > current_tracer
echo 1 > tracing_on
```

