
  function callback() {
    var str = '{"artist":"So and So","title":"Not Relevant"}';
    var jsonObj = JSON.parse(str);
  
  $("#d").append('<ul><li><a onclick="alert(JSON.stringify(jsonObj))" title="NA">Link</a></li></ul>');
  };

