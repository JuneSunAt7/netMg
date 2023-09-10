
  function callback() {

    let file = "tree.json";;

    let reader = new FileReader();
  
    reader.readAsText(file);
  
    reader.onload = function() {
      console.log(reader.result);
    };
  
    reader.onerror = function() {
      console.log(reader.error);
    };
  };