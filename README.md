# beehive do not need to handle config.

the core function of beehive is handling the cross-module communication. but it also handle config (such as commandline config and yaml config).

i think beehive don`t need to handle those config. we need let kubeedge to handle it .

currently, edgemesh , keadm  process their own profiles individually. howerver edgecontroller, edge_core,edgesite use beehive config.

they all lack of uniform  handling.

beehive's config has many drawbacks。

When the program is running, it will output a lot of logs first.

this is an example:

if we want edge_core output current detail version. it`s going to be very difficult.

because edge_core first import github.com/kubeedge/beehive/pkg/core, and "github.com/kubeedge/beehive/pkg/core" will import "github.com/kubeedge/beehive/pkg/common/config". and they will output many log info rather than only output current version.

I recommend removing the beehive config processing method.  Use a unified config handle, like keadm



111
