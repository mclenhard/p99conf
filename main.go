package main

import (
	"context"
	"github.com/fntlnz/mountinfo"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

)

func main(){
     podName := os.Args[1]
     namespace := os.Args[2]
     configuration, _ := rest.InClusterConfig()
	 client, _ := kubernetes.NewForConfig(configuration)
     pod,_ := getPod(podName,namespace,client)
	 pids,_ := GetProcessId(string(pod.UID))
	 log.Printf("pids %v",pids)

}
func getPod(name string ,namespace string,client *kubernetes.Clientset)(*apiv1.Pod,error){
	pod,err := client.CoreV1().Pods(namespace).Get(context.TODO(),name,metav1.GetOptions{})
	if err != nil{
		return nil,err
	}
	return pod,nil
}


func GetProcessId(uid string) ([]int,error){
	directory,err := os.Open("/host/proc")
	log.Printf("looking for uid %s",uid)
	if err != nil{
		log.Println(err)
	}
	var pids []int
	defer directory.Close()
	for {
		dir, err := directory.ReadDir(0)
		if err != nil {
			return pids, err
		}
		for _, d := range dir {
			if _, err := strconv.ParseInt(d.Name(),10,64); err != nil {
				continue
			}

			mountinfoPath := path.Join("/host/proc",d.Name(),"mountinfo")
			mount, err := mountinfo.GetMountInfo(mountinfoPath)
			if err != nil{
				return pids,err
			}
			for _,m := range mount{{
					if(strings.Contains(m.Root,uid)){
						pid, err := strconv.Atoi(d.Name())
						if err != nil {
							return pids,nil
						}
						pids = append(pids,pid)
						break
					}
				}
			}
		}
		return pids, err

	}
}

