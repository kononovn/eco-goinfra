package main

import (
	"log"
	"os"

	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"github.com/openshift-kni/eco-goinfra/pkg/metallb"
	"github.com/openshift-kni/eco-goinfra/pkg/namespace"
)

func main() {
	log.Print("HELLO")

	clinet := clients.New("")

	if clinet == nil {
		log.Print("clinet is nill")
		os.Exit(1)
	}

	nsbuilder := namespace.NewBuilder(clinet, "mynamespace")
	nsbuild, err := nsbuilder.Create()

	if err != nil {
		log.Print("Fail to create ns")
		os.Exit(1)
	}

	err = nsbuild.Delete()
	if err != nil {
		log.Print("Fail to delete ns")
		os.Exit(1)
	}

	metallbio, err := metallb.PullAddressPool(clinet, "peer-one", "metallb-system")
	log.Print(err)
	log.Print(metallbio.Exists())
	metallbio, err = metallbio.Delete()
	//metallbio, err = metallbio.Create()
	//_, err = metallbio.Create()
	//log.Print(metallbio.Delete())
	log.Print(err)
	log.Print(metallbio.Definition.Spec)
}
