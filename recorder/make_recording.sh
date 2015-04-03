#! /bin/bash
if [[ -z "$RS_KEY" ]]; then
  echo "Please set RS_KEY to your API key"
  exit 1
fi

rm -f recording-new.json
ARGS=(--host us-3.rightscale.com --key $RS_KEY)

set -x

./recorder ${ARGS[@]} cm15 index clouds
./recorder ${ARGS[@]} cm15 index /api/clouds
./recorder ${ARGS[@]} cm15 show /api/clouds/6

./recorder ${ARGS[@]} --x1 ".cloud_type" cm15 show /api/clouds/6
./recorder ${ARGS[@]} --xm ".cloud_type" cm15 index clouds
./recorder ${ARGS[@]} --xj ".cloud_type" cm15 index clouds

./recorder ${ARGS[@]} --x1 '*:has(.name:val("EC2 us-east-1")) .name' cm15 index clouds
./recorder ${ARGS[@]} --xm '*:has(.name:val("EC2 us-east-1")) .name' cm15 index clouds
./recorder ${ARGS[@]} --xj '*:has(.name:val("EC2 us-east-1")) .name' cm15 index clouds

./recorder ${ARGS[@]} --xm .local_disks cm15 index /api/clouds/1/instance_types

./recorder ${ARGS[@]} --x1 :root cm15 index /api/clouds/3/volume_snapshots \
  'filter[]=resource_uid==snap-00828462'
./recorder ${ARGS[@]} --xm :root cm15 index /api/clouds/3/volume_snapshots \
  'filter[]=resource_uid==snap-00828462'
./recorder ${ARGS[@]} --xj :root cm15 index /api/clouds/3/volume_snapshots \
  'filter[]=resource_uid==snap-00828462'

names=`./recorder ${ARGS[@]} --xm '*:has(.cloud_type:val("amazon")) .name' cm15 index clouds`
declare -a "names=($names)" # magic
set | egrep '^names'
[[ ${#names[@]} = 9 ]] || exit 1
./recorder ${ARGS[@]} --xj '*:has(.cloud_type:val("amazon")) .name' cm15 index clouds

./recorder ${ARGS[@]} --x1 'object:has(.name:val("rsc-test"))' cm15 index deployments
href=`./recorder ${ARGS[@]} --xh location create deployments \
  'deployment[name]=rsc-test' \
  'deployment[description]=expendable deployment used to test rsc'`
declare -a "href=($href)"
./recorder ${ARGS[@]} cm15 destroy $href

# find existing deployment and delete it
deployment=`./recorder ${ARGS[@]} \
  --x1 ':has(.rel:val("self")).href' \
  index deployments \
  'filter[]=name==rsc-test'`
echo "deployment: $deployment"
if [[ -n "$deployment" ]]; then
  ./recorder ${ARGS[@]} cm15 destroy $deployment
fi

# create a deployment with too few params -> error
./recorder ${ARGS[@]} --xh location cm15 create deployments

# create a deployment to launch an instance in
deployment_href=`./recorder ${ARGS[@]} --xh location cm15 create deployments 'deployment[name]=rsc-test'`
echo "deployment_href: $deployment_href"

# get the href of the EC2 us-east-1 cloud
cloud_href=`./recorder ${ARGS[@]} \
  --x1 '*:has(.name:val("EC2 us-east-1")) :has(.rel:val("self")).href' \
   cm15 index clouds`
echo "cloud_href: $cloud_href"

# locate the image to launch
image_href=`./recorder ${ARGS[@]} \
  --x1 ':has(.rel:val("self")).href' \
  cm15 index $cloud_href/images \
  'filter[]=resource_uid==ami-6089d208'`
echo "image_href: $image_href"

# locate an instance type
inst_type_href=`./recorder ${ARGS[@]} \
  --x1 ':has(.rel:val("self")).href' \
  cm15 index $cloud_href/instance_types \
  'filter[]=name==m3.medium'`
echo "inst_type_href: $inst_type_href"

# launch the instance
instance_href=`./recorder ${ARGS[@]} \
  --xh location\
  cm15 create $cloud_href/instances \
  "instance[image_href]=$image_href" \
  "instance[instance_type_href]=$inst_type_href" \
  "instance[name]=rsc-test"`
echo "instance_href: $instance_href"

# wait for it to be running
while true; do
  state=`./recorder ${ARGS[@]} \
    --xm '.state' \
    cm15 show $instance_href`
  echo state: $state
  if [[ "$state" != pending ]]; then break; fi
  sleep 60
done

./recorder ${ARGS[@]} cm15 show $instance_href
./recorder ${ARGS[@]} --x1 .locked cm15 show $instance_href
./recorder ${ARGS[@]} --xm .locked cm15 show $instance_href
./recorder ${ARGS[@]} --xj .locked cm15 show $instance_href

# find server template
st_href=`./recorder ${ARGS[@]} \
  --x1 ':has(.rel:val("self")).href' \
  cm15 index server_templates \
  'filter[]=name==Rightlink 10.0.rc4 Linux Base' \
  'filter[]=revision==0'`
echo st_href: $st_href

# terminating instance
./recorder ${ARGS[@]} cm15 terminate $instance_href
./recorder ${ARGS[@]} cm15 destroy $deployment_href



