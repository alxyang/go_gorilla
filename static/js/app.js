$( document ).ready(function() {
  console.log( "document loaded" );

  $.get( "/user/777", function( data ) {
    console.log("sample get request passed");
  });

  // Assign handlers immediately after making the request,
  // and remember the jqxhr object for this request
  var sampleData =   {
    name:"Donald Duck",
    city:"Duckburg"
  };
  var jqxhr = $.post( "/test", sampleData, function() {
    console.log( "successfully posted data" );
  })
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

