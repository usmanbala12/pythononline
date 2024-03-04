var editor;

document.addEventListener("DOMContentLoaded", (event) => {
    editor = ace.edit("editor");
    editor.setTheme("ace/theme/monokai");
    editor.getSession().setMode("ace/mode/python");
});


function outf(text) {
    var mypre = document.getElementById("output");
    mypre.innerHTML = mypre.innerHTML + text;
}

function builtinRead(x) {
    if (Sk.builtinFiles === undefined || Sk.builtinFiles["files"][x] === undefined)
        throw "File not found: '" + x + "'";
    return Sk.builtinFiles["files"][x];
}

function runit(event) {

    if (window.innerWidth <= 600) {
        const runner = document.querySelector('.runner')
        const visibility = runner.getAttribute("data-visible")
        if (visibility === "false") {
            runner.setAttribute("data-visible", true)
            event.target.innerHTML = "Close"
        } else {
            runner.setAttribute("data-visible", false)
            event.target.innerHTML = "Run"
        }
    }


    var prog = editor.getValue()
    var mypre = document.getElementById("output");
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
            document.getElementById("output").style.color = "red"
            outf(err.toString());
        });
}