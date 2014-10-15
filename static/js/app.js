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
      appendLog($("<div><b>Connection closed.</b></div>"));
    };

    conn.onerror = function(error){
      console.log('Error detected: ' + JSON.stringify(error));
    };

    conn.onmessage = function(evt) {
      // console.log(evt.data);
      appendLog($("<div/>").text(evt.data));
    };

  } else {
    console.log( "browser does support websockets" );
  }

    var msg = $("#msg");
    var log = $("#log");

    function appendLog(msg) {
        var d = log[0]
        var doScroll = d.scrollTop == d.scrollHeight - d.clientHeight;
        msg.appendTo(log)
        if (doScroll) {
            d.scrollTop = d.scrollHeight - d.clientHeight;
        }
    }

    $("#form").submit(function() {
        if (!conn) {
            return false;
        }
        if (!msg.val()) {
            return false;
        }
        conn.send(msg.val());
        msg.val("");
        return false
    });



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

