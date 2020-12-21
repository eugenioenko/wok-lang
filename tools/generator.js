let fs = require('fs');

const ExpressionAST = {
    Assign: ['name Token', 'value Expression'],
    Binary: ['left Expression', 'operator Token', 'right Expression'],
    Call: ['callee Expression', 'paren Token', 'args []Expression'],
    Grouping: ['expression Expression'],
    Literal: ['value DataType'],
    Unary: ['operator Token', 'right Expression'],
    Value: ['value Token'],
    Variable: ['value Token']
};

const StatementAST = {

};

function generateAST(base, arg, AST, filename) {
    let file =
`package main

type Expression interface {
    Accept(visitor VisitorExpression) DataType
}
\n`;

    file += `type Visitor${base} interface {\n`;
    Object.keys(AST).forEach(name => {
        file += `\tVisit${base}${name}(${arg} Expression${name}) DataType\n`;
    });
    file += '}\n\n';

    Object.keys(AST).forEach(name => {
        const syntax = AST[name];
        file += `type ${base}${name} struct {\n`;
        syntax.forEach(member => {
            file += '    ' + member + '\n'
        });
        file += '}\n'
        file += `\nfunc Make${base}${name}(${syntax.join(', ')}) ${base}${name} {\n`
        file += `\treturn ${base}${name}{`
        file += syntax.map(member => member.split(' ')[0]).join(', ')
        file += '}\n}\n'
        file += `\nfunc (${arg} ${base}${name}) Accept (visitor Visitor${base}) DataType {\n`
        file += `\treturn visitor.Visit${base}${name}(${arg})\n}\n\n`
    });

    fs.writeFile(`${filename}.go`, file, function (err, data) {
        if (err) console.log(err);
        console.log(`${filename}.go generated`);
    });
}

generateAST('Expression', 'expr', ExpressionAST, 'expressions');