function main(){
  getGPS();
}

function getGPS(){
  var wOptions = {
    "enableHighAccuracy": true,                       // true : 高精度
    "timeout": 1000,                                 // タイムアウト : ミリ秒
    "maximumAge": 0                                 // データをキャッシュ時間 : ミリ秒
  };
  var id=navigator.geolocation.watchPosition(getGPSsuccess,getGPSFailure,wOptions);

}

function getGPSsuccess(pos){
  var crd = pos.coords;
  document.getElementById("latitude").innerHTML+="lat:"+crd.latitude+"\n";
  document.getElementById("longitude").innerHTML+="lon:"+crd.longitude+"\n";
}
function getGPSFailure(args){
  console.log("failure");
  console.log(args)
}

main()
