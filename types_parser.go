package typesparser

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

func Plugin(pb *parser.Builder) {
	pb.UseStatementInterceptor(func(p *parser.Parser, next func() ast.Statement) ast.Statement {
		if p.CurrentToken.Type != token.LET {
			return next()
		}

		stmt := &ast.LetStatement{Token: p.CurrentToken}
		if !p.ExpectToken(token.IDENT) {
			return nil
		}

		// simply ignore the annotation types :)
		if p.PeekToken.Type == token.COLON {
			p.NextToken() // consume :
			if !p.ExpectToken(token.IDENT) {
				return nil
			}
		}

		stmt.Name = &ast.Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
		if p.PeekToken.Type == token.ASSIGN {
			p.NextToken() // consume =
			p.NextToken() // move to value
			stmt.Value = p.ParseExpression()
		}
		if p.PeekToken.Type == token.SEMICOLON {
			p.NextToken()
		}
		return stmt
	})
}
