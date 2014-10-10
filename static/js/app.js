$( document ).ready(function() {
  console.log( "document loaded" );

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

