package main

import (
	"fmt"
	"reflect"
	"regexp"

	pb "github.com/dundee/gogenerate/protobuf"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var msgs = []proto.Message{
	&pb.User{},
}

func main() {
	nameReg := regexp.MustCompile(`name=([A-Za-z0-9_]+),`)

	for _, msg := range msgs {
		logFields := map[string]string{}

		t := reflect.TypeOf(msg).Elem()
		desc := msg.ProtoReflect().Descriptor()

		descFields := desc.Fields()

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			tag := f.Tag.Get("protobuf")
			if tag == "" {
				continue
			}

			name := nameReg.FindStringSubmatch(tag)[1]

			descf := descFields.ByName(protoreflect.Name(name))
			if proto.HasExtension(descf.Options(), pb.E_Key) {
				ex := proto.GetExtension(descf.Options(), pb.E_Key)
				if logKey, ok := ex.(string); ok {
					logFields[logKey] = f.Name
				}
			}
		}

		fmt.Printf("%s: %v\n", t.Name(), logFields)
	}

}
