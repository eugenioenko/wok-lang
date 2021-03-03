let fs = require('fs');

const ExpressionAST = {
    Assign: ['name Expression', 'value Expression'],
    Binary: ['left Expression', 'operator Token', 'right Expression'],
    Call: ['callee Expression', 'paren Token', 'args []Expression'],
    Grouping: ['expression Expression'],
    Literal: ['value WokData'],
    Unary: ['operator Token', 'right Expression'],
    Value: ['value Token'],
    Variable: ['value Token']
};

const StatementAST = {
    Expression : ['expr Expression'],
    Func: ['name Token', 'params []Token', 'body []Statement'],
    If: ['condition Expression', 'thenStmt Statement', 'elseStmt Statement'],
    Return: ['keyword Token', 'value Expression'],
    Var: ['name Token', 'dtype Token', 'initial Expression', 'writable bool'],
    While: ['condition Expression', 'loop Statement'],
    Print: ['value Expression']
};


function generateAST(base, arg, AST, filename) {
    let file =
`package main

type ${base} interface {
    Accept(visitor Visitor${base}) WokData
}\n\n`;

    file += `type Visitor${base} interface {\n`;
    Object.keys(AST).forEach(name => {
        file += `\tVisit${base}${name}(${arg} *${base}${name}) WokData\n`;
    });
    file += '}\n\n';

    Object.keys(AST).forEach(name => {
        const syntax = AST[name];
        file += `type ${base}${name} struct {\n`;
        syntax.forEach(member => {
            file += '    ' + member + '\n'
        });
        file += '}\n'
        file += `\nfunc New${base}${name}(${syntax.join(', ')}) *${base}${name} {\n`
        file += `\treturn &${base}${name}{`
        file += syntax.map(member => member.split(' ')[0]).join(', ')
        file += '}\n}\n'
        file += `\nfunc (${arg} *${base}${name}) Accept (visitor Visitor${base}) WokData {\n`
        file += `\treturn visitor.Visit${base}${name}(${arg})\n}\n\n`
    });

    fs.writeFile(`${filename}.go`, file, function (err, data) {
        if (err) console.log(err);
        console.log(`${filename}.go generated`);
    });
}

generateAST('Expression', 'expr', ExpressionAST, 'expressions');
generateAST('Statement', 'stmt', StatementAST, 'statements');