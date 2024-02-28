var editorsArray = [];

$(document).ready(function () {

    var editor;
    var ednum = 0;

    $('.editor').each(function (index) {
        ednum++;
        _name = "editor_" + ednum;
        code = "var name = $(this).data('name');"
            + _name + " = ace.edit(this);"
            + _name + ".setTheme('ace/theme/monokai');"
            + _name + ".getSession().setMode('ace/mode/python');"
            + "var code_" + ednum + " = $('textarea[name='+name+']');"
        eval(code);

        editorsArray.push(window[_name]);
    });

});


function runExample(btn) {
    let dataNo = btn.getAttribute("data-no")

    function outf(text) {
        var mypre = document.getElementById("output" + dataNo);
        mypre.innerHTML = mypre.innerHTML + text;
    }

    function builtinRead(x) {
        if (Sk.builtinFiles === undefined || Sk.builtinFiles["files"][x] === undefined)
            throw "File not found: '" + x + "'";
        return Sk.builtinFiles["files"][x];
    }

    let runner = document.getElementById("runner" + dataNo)

    if (getComputedStyle(runner).display === "none") {
        btn.innerHTML = "Close"
        runner.style.display = "block"
        document.getElementById("editor_" + dataNo).style.display = "none"

        var prog = editorsArray[parseInt(dataNo) - 1].getValue()
        var mypre = document.getElementById("output" + dataNo);
        mypre.style.color = "white"
        mypre.innerHTML = '';
        Sk.pre = "output";
        Sk.configure({ output: outf, read: builtinRead });
        (Sk.TurtleGraphics || (Sk.TurtleGraphics = {})).target = 'mycanvas';
        var myPromise = Sk.misceval.asyncToPromise(function () {
            return Sk.importMainWithBody("<stdin>", false, prog, true);
        });
        myPromise.then(function (mod) {
            console.log('success');
        },
            function (err) {
                document.getElementById("output" + dataNo).style.color = "red"
                outf(err.toString());
            });
    }

    else {
        runner.style.display = "none"
        document.getElementById("editor_" + dataNo).style.display = "block"
        btn.innerHTML = "Run code"
    }
}