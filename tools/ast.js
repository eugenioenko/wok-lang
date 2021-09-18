let fs = require("fs");

const ExpressionAST = {
  List: ["value []Expression"],
  Atom: ["value Token"],
};

function generateAST(base, arg, AST, filename) {
  let file = `package main

type ${base} interface {
    Accept(visitor Visitor${base}) WokData
}\n\n`;

  file += `type Visitor${base} interface {\n`;
  Object.keys(AST).forEach((name) => {
    file += `\tVisit${base}${name}(${arg} *${base}${name}) WokData\n`;
  });
  file += "}\n\n";

  Object.keys(AST).forEach((name) => {
    const syntax = AST[name];
    file += `type ${base}${name} struct {\n`;
    syntax.forEach((member) => {
      file += "    " + member + "\n";
    });
    file += "}\n";
    file += `\nfunc New${base}${name}(${syntax.join(
      ", "
    )}) *${base}${name} {\n`;
    file += `\treturn &${base}${name}{`;
    file += syntax.map((member) => member.split(" ")[0]).join(", ");
    file += "}\n}\n";
    file += `\nfunc (${arg} *${base}${name}) Accept (visitor Visitor${base}) WokData {\n`;
    file += `\treturn visitor.Visit${base}${name}(${arg})\n}\n\n`;
  });

  fs.writeFile(`${filename}.go`, file, function (err, data) {
    if (err) console.log(err);
    console.log(`${filename}.go generated`);
  });
}

generateAST("Expression", "expr", ExpressionAST, "expressions");
