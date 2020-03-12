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
  console.log(crd.latitude);
  console.log(crd.longitude);
}
function getGPSFailure(args){
  console.log("failure");
  console.log(args)
}

main()
