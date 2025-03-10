# checking the fond size

# - detection
# - 

# write monitoring script which reads a file and file contains a list of devices 
# check the reachability of those devices and tell me which devices are not reachable 

import subprocess
import threading

def read_devices(file_path):
    with open(file_path,'r') as file:
        devices = [line.strip() for line in file if line.strip()]

def is_reachable(host, unreachable_devices):
    try:
        result = subprocess.run(["ping","-c","1","-W","1", host]), stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        if result.returncode !=0:
            unreachable_devices.append(host)
    except Exception as e:
        print(f"Error ping {host}: {e}")
        unreachable_devices.append(host)
    
def check_reachability(file_path, batch_size=5):
    devices = read_devices(file_path)
    unreachable_devices = []
    threads = []


    for i in range(0, len(devices),batch_size):
        j = i+batch_size
        if j> len(devices)
        for device in devices[i:i+batch_size]:
            # if not is_reachable(devices):
            #     unreachable_devices.append(devices)
            thread = threading.Thread(target=is_reachable, args=(device,unreachable_devices))
            threads.append(thread)
            thread.start()
    
        for thread in threads:
            thread.join()
    
    if unreachable_devices:
        print("Unreachable devices:")
        for device in unreachable_devices:
            print(device)
    else:
        print("All devices are reachable.")
    
if __name__=="__main__":
    file_path = "devices.txt"
    check_reachability(file_path)

        
