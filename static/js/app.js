$(function() {

  var conn;

  if (window.WebSocket) {
    console.log("websockets available");
    conn = new WebSocket("ws://localhost:8080/ws");

    conn.onopen = function(){
      /*Send a small message to the console once the connection is established */
      console.log('Connection open!');
    };

    conn.onclose = function(){
      console.log('Connection closed');
    };

    conn.onerror = function(error){
      console.log('Error detected: ' + JSON.stringify(error));
    };

  } else {
    console.log( "browser does support websockets" );
  }

  //sample GET request to server
  $.get( "/user/777", function( data ) {
    console.log("sample get request passed");
  });

  // Assign handlers immediately after making the request,
  // and remember the jqxhr object for this request
  var sampleData =   {
    "name":"Donald Duck",
    "city":"Duckburg"
  };

  //sample POST request to server
  var jqxhr = $.post( "/test", JSON.stringify(sampleData), function() {
    console.log( "successfully posted data" );
  }, "json")
    .done(function() {
      console.log( "second success" );
    })
    .fail(function() {
      console.log( "error" );
    })
    .always(function() {
      console.log( "finished" );
  });

});

