package typesparser

import (
	"github.com/xjslang/xjs/ast"
	"github.com/xjslang/xjs/parser"
	"github.com/xjslang/xjs/token"
)

func ParseFunctionParameters(p *parser.Parser) []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	if p.PeekToken.Type == token.RPAREN {
		p.NextToken()
		return identifiers
	}
	p.NextToken()
	ident := &ast.Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
	identifiers = append(identifiers, ident)

	// simply ignore the annotation types :)
	if p.PeekToken.Type == token.COLON {
		p.NextToken() // consume :
		if !p.ExpectToken(token.IDENT) {
			return nil
		}
	}

	for p.PeekToken.Type == token.COMMA {
		p.NextToken()
		p.NextToken()
		ident := &ast.Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
		identifiers = append(identifiers, ident)

		// simply ignore the annotation types :)
		if p.PeekToken.Type == token.COLON {
			p.NextToken() // consume :
			if !p.ExpectToken(token.IDENT) {
				return nil
			}
		}
	}
	if !p.ExpectToken(token.RPAREN) {
		return nil
	}
	return identifiers
}

func Plugin(pb *parser.Builder) {
	// intercepts 'let' statements
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

	// intercepts 'function' statements
	pb.UseStatementInterceptor(func(p *parser.Parser, next func() ast.Statement) ast.Statement {
		if p.CurrentToken.Type != token.FUNCTION {
			return next()
		}

		stmt := &ast.FunctionDeclaration{Token: p.CurrentToken}
		if !p.ExpectToken(token.IDENT) {
			return nil
		}
		stmt.Name = &ast.Identifier{Token: p.CurrentToken, Value: p.CurrentToken.Literal}
		if !p.ExpectToken(token.LPAREN) {
			return nil
		}
		stmt.Parameters = ParseFunctionParameters(p)
		if !p.ExpectToken(token.LBRACE) {
			return nil
		}
		p.PushContext(parser.FunctionContext)
		defer p.PopContext()
		stmt.Body = p.ParseBlockStatement()
		return stmt
	})
}
