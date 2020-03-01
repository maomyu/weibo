

function searchWeibo() {
    var text=document.getElementById("searchText").value;
    var searchurl = "/searchUser?showtype=search&searchName=" + text;
    console.log("url: "+searchurl);
    window.location.href= searchurl;
}
