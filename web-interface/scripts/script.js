
  function callback() {
    var str = '{"artist":"So and So","title":"Not Relevant"}';
    var jsonObj = JSON.parse(str);

    $("#d").append('<a onclick="console.log(JSON.stringify(jsonObj))" >Link</a>');

  };

