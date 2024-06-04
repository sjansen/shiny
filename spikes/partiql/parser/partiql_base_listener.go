// Code generated from PartiQL.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // PartiQL
import "github.com/antlr4-go/antlr/v4"

// BasePartiQLListener is a complete listener for a parse tree produced by PartiQLParser.
type BasePartiQLListener struct{}

var _ PartiQLListener = &BasePartiQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePartiQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePartiQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePartiQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePartiQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BasePartiQLListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasePartiQLListener) ExitRoot(ctx *RootContext) {}

// EnterQueryDql is called when production QueryDql is entered.
func (s *BasePartiQLListener) EnterQueryDql(ctx *QueryDqlContext) {}

// ExitQueryDql is called when production QueryDql is exited.
func (s *BasePartiQLListener) ExitQueryDql(ctx *QueryDqlContext) {}

// EnterQueryDml is called when production QueryDml is entered.
func (s *BasePartiQLListener) EnterQueryDml(ctx *QueryDmlContext) {}

// ExitQueryDml is called when production QueryDml is exited.
func (s *BasePartiQLListener) ExitQueryDml(ctx *QueryDmlContext) {}

// EnterQueryDdl is called when production QueryDdl is entered.
func (s *BasePartiQLListener) EnterQueryDdl(ctx *QueryDdlContext) {}

// ExitQueryDdl is called when production QueryDdl is exited.
func (s *BasePartiQLListener) ExitQueryDdl(ctx *QueryDdlContext) {}

// EnterQueryExec is called when production QueryExec is entered.
func (s *BasePartiQLListener) EnterQueryExec(ctx *QueryExecContext) {}

// ExitQueryExec is called when production QueryExec is exited.
func (s *BasePartiQLListener) ExitQueryExec(ctx *QueryExecContext) {}

// EnterExplainOption is called when production explainOption is entered.
func (s *BasePartiQLListener) EnterExplainOption(ctx *ExplainOptionContext) {}

// ExitExplainOption is called when production explainOption is exited.
func (s *BasePartiQLListener) ExitExplainOption(ctx *ExplainOptionContext) {}

// EnterAsIdent is called when production asIdent is entered.
func (s *BasePartiQLListener) EnterAsIdent(ctx *AsIdentContext) {}

// ExitAsIdent is called when production asIdent is exited.
func (s *BasePartiQLListener) ExitAsIdent(ctx *AsIdentContext) {}

// EnterAtIdent is called when production atIdent is entered.
func (s *BasePartiQLListener) EnterAtIdent(ctx *AtIdentContext) {}

// ExitAtIdent is called when production atIdent is exited.
func (s *BasePartiQLListener) ExitAtIdent(ctx *AtIdentContext) {}

// EnterByIdent is called when production byIdent is entered.
func (s *BasePartiQLListener) EnterByIdent(ctx *ByIdentContext) {}

// ExitByIdent is called when production byIdent is exited.
func (s *BasePartiQLListener) ExitByIdent(ctx *ByIdentContext) {}

// EnterSymbolPrimitive is called when production symbolPrimitive is entered.
func (s *BasePartiQLListener) EnterSymbolPrimitive(ctx *SymbolPrimitiveContext) {}

// ExitSymbolPrimitive is called when production symbolPrimitive is exited.
func (s *BasePartiQLListener) ExitSymbolPrimitive(ctx *SymbolPrimitiveContext) {}

// EnterDql is called when production dql is entered.
func (s *BasePartiQLListener) EnterDql(ctx *DqlContext) {}

// ExitDql is called when production dql is exited.
func (s *BasePartiQLListener) ExitDql(ctx *DqlContext) {}

// EnterExecCommand is called when production execCommand is entered.
func (s *BasePartiQLListener) EnterExecCommand(ctx *ExecCommandContext) {}

// ExitExecCommand is called when production execCommand is exited.
func (s *BasePartiQLListener) ExitExecCommand(ctx *ExecCommandContext) {}

// EnterDdl is called when production ddl is entered.
func (s *BasePartiQLListener) EnterDdl(ctx *DdlContext) {}

// ExitDdl is called when production ddl is exited.
func (s *BasePartiQLListener) ExitDdl(ctx *DdlContext) {}

// EnterCreateTable is called when production CreateTable is entered.
func (s *BasePartiQLListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production CreateTable is exited.
func (s *BasePartiQLListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateIndex is called when production CreateIndex is entered.
func (s *BasePartiQLListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production CreateIndex is exited.
func (s *BasePartiQLListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterDropTable is called when production DropTable is entered.
func (s *BasePartiQLListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production DropTable is exited.
func (s *BasePartiQLListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropIndex is called when production DropIndex is entered.
func (s *BasePartiQLListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production DropIndex is exited.
func (s *BasePartiQLListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterDmlBaseWrapper is called when production DmlBaseWrapper is entered.
func (s *BasePartiQLListener) EnterDmlBaseWrapper(ctx *DmlBaseWrapperContext) {}

// ExitDmlBaseWrapper is called when production DmlBaseWrapper is exited.
func (s *BasePartiQLListener) ExitDmlBaseWrapper(ctx *DmlBaseWrapperContext) {}

// EnterDmlDelete is called when production DmlDelete is entered.
func (s *BasePartiQLListener) EnterDmlDelete(ctx *DmlDeleteContext) {}

// ExitDmlDelete is called when production DmlDelete is exited.
func (s *BasePartiQLListener) ExitDmlDelete(ctx *DmlDeleteContext) {}

// EnterDmlInsertReturning is called when production DmlInsertReturning is entered.
func (s *BasePartiQLListener) EnterDmlInsertReturning(ctx *DmlInsertReturningContext) {}

// ExitDmlInsertReturning is called when production DmlInsertReturning is exited.
func (s *BasePartiQLListener) ExitDmlInsertReturning(ctx *DmlInsertReturningContext) {}

// EnterDmlBase is called when production DmlBase is entered.
func (s *BasePartiQLListener) EnterDmlBase(ctx *DmlBaseContext) {}

// ExitDmlBase is called when production DmlBase is exited.
func (s *BasePartiQLListener) ExitDmlBase(ctx *DmlBaseContext) {}

// EnterDmlBaseCommand is called when production dmlBaseCommand is entered.
func (s *BasePartiQLListener) EnterDmlBaseCommand(ctx *DmlBaseCommandContext) {}

// ExitDmlBaseCommand is called when production dmlBaseCommand is exited.
func (s *BasePartiQLListener) ExitDmlBaseCommand(ctx *DmlBaseCommandContext) {}

// EnterPathSimple is called when production pathSimple is entered.
func (s *BasePartiQLListener) EnterPathSimple(ctx *PathSimpleContext) {}

// ExitPathSimple is called when production pathSimple is exited.
func (s *BasePartiQLListener) ExitPathSimple(ctx *PathSimpleContext) {}

// EnterPathSimpleLiteral is called when production PathSimpleLiteral is entered.
func (s *BasePartiQLListener) EnterPathSimpleLiteral(ctx *PathSimpleLiteralContext) {}

// ExitPathSimpleLiteral is called when production PathSimpleLiteral is exited.
func (s *BasePartiQLListener) ExitPathSimpleLiteral(ctx *PathSimpleLiteralContext) {}

// EnterPathSimpleSymbol is called when production PathSimpleSymbol is entered.
func (s *BasePartiQLListener) EnterPathSimpleSymbol(ctx *PathSimpleSymbolContext) {}

// ExitPathSimpleSymbol is called when production PathSimpleSymbol is exited.
func (s *BasePartiQLListener) ExitPathSimpleSymbol(ctx *PathSimpleSymbolContext) {}

// EnterPathSimpleDotSymbol is called when production PathSimpleDotSymbol is entered.
func (s *BasePartiQLListener) EnterPathSimpleDotSymbol(ctx *PathSimpleDotSymbolContext) {}

// ExitPathSimpleDotSymbol is called when production PathSimpleDotSymbol is exited.
func (s *BasePartiQLListener) ExitPathSimpleDotSymbol(ctx *PathSimpleDotSymbolContext) {}

// EnterReplaceCommand is called when production replaceCommand is entered.
func (s *BasePartiQLListener) EnterReplaceCommand(ctx *ReplaceCommandContext) {}

// ExitReplaceCommand is called when production replaceCommand is exited.
func (s *BasePartiQLListener) ExitReplaceCommand(ctx *ReplaceCommandContext) {}

// EnterUpsertCommand is called when production upsertCommand is entered.
func (s *BasePartiQLListener) EnterUpsertCommand(ctx *UpsertCommandContext) {}

// ExitUpsertCommand is called when production upsertCommand is exited.
func (s *BasePartiQLListener) ExitUpsertCommand(ctx *UpsertCommandContext) {}

// EnterRemoveCommand is called when production removeCommand is entered.
func (s *BasePartiQLListener) EnterRemoveCommand(ctx *RemoveCommandContext) {}

// ExitRemoveCommand is called when production removeCommand is exited.
func (s *BasePartiQLListener) ExitRemoveCommand(ctx *RemoveCommandContext) {}

// EnterInsertCommandReturning is called when production insertCommandReturning is entered.
func (s *BasePartiQLListener) EnterInsertCommandReturning(ctx *InsertCommandReturningContext) {}

// ExitInsertCommandReturning is called when production insertCommandReturning is exited.
func (s *BasePartiQLListener) ExitInsertCommandReturning(ctx *InsertCommandReturningContext) {}

// EnterInsertLegacy is called when production InsertLegacy is entered.
func (s *BasePartiQLListener) EnterInsertLegacy(ctx *InsertLegacyContext) {}

// ExitInsertLegacy is called when production InsertLegacy is exited.
func (s *BasePartiQLListener) ExitInsertLegacy(ctx *InsertLegacyContext) {}

// EnterInsert is called when production Insert is entered.
func (s *BasePartiQLListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production Insert is exited.
func (s *BasePartiQLListener) ExitInsert(ctx *InsertContext) {}

// EnterOnConflictLegacy is called when production OnConflictLegacy is entered.
func (s *BasePartiQLListener) EnterOnConflictLegacy(ctx *OnConflictLegacyContext) {}

// ExitOnConflictLegacy is called when production OnConflictLegacy is exited.
func (s *BasePartiQLListener) ExitOnConflictLegacy(ctx *OnConflictLegacyContext) {}

// EnterOnConflict is called when production OnConflict is entered.
func (s *BasePartiQLListener) EnterOnConflict(ctx *OnConflictContext) {}

// ExitOnConflict is called when production OnConflict is exited.
func (s *BasePartiQLListener) ExitOnConflict(ctx *OnConflictContext) {}

// EnterConflictTarget is called when production conflictTarget is entered.
func (s *BasePartiQLListener) EnterConflictTarget(ctx *ConflictTargetContext) {}

// ExitConflictTarget is called when production conflictTarget is exited.
func (s *BasePartiQLListener) ExitConflictTarget(ctx *ConflictTargetContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BasePartiQLListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BasePartiQLListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterConflictAction is called when production conflictAction is entered.
func (s *BasePartiQLListener) EnterConflictAction(ctx *ConflictActionContext) {}

// ExitConflictAction is called when production conflictAction is exited.
func (s *BasePartiQLListener) ExitConflictAction(ctx *ConflictActionContext) {}

// EnterDoReplace is called when production doReplace is entered.
func (s *BasePartiQLListener) EnterDoReplace(ctx *DoReplaceContext) {}

// ExitDoReplace is called when production doReplace is exited.
func (s *BasePartiQLListener) ExitDoReplace(ctx *DoReplaceContext) {}

// EnterDoUpdate is called when production doUpdate is entered.
func (s *BasePartiQLListener) EnterDoUpdate(ctx *DoUpdateContext) {}

// ExitDoUpdate is called when production doUpdate is exited.
func (s *BasePartiQLListener) ExitDoUpdate(ctx *DoUpdateContext) {}

// EnterUpdateClause is called when production updateClause is entered.
func (s *BasePartiQLListener) EnterUpdateClause(ctx *UpdateClauseContext) {}

// ExitUpdateClause is called when production updateClause is exited.
func (s *BasePartiQLListener) ExitUpdateClause(ctx *UpdateClauseContext) {}

// EnterSetCommand is called when production setCommand is entered.
func (s *BasePartiQLListener) EnterSetCommand(ctx *SetCommandContext) {}

// ExitSetCommand is called when production setCommand is exited.
func (s *BasePartiQLListener) ExitSetCommand(ctx *SetCommandContext) {}

// EnterSetAssignment is called when production setAssignment is entered.
func (s *BasePartiQLListener) EnterSetAssignment(ctx *SetAssignmentContext) {}

// ExitSetAssignment is called when production setAssignment is exited.
func (s *BasePartiQLListener) ExitSetAssignment(ctx *SetAssignmentContext) {}

// EnterDeleteCommand is called when production deleteCommand is entered.
func (s *BasePartiQLListener) EnterDeleteCommand(ctx *DeleteCommandContext) {}

// ExitDeleteCommand is called when production deleteCommand is exited.
func (s *BasePartiQLListener) ExitDeleteCommand(ctx *DeleteCommandContext) {}

// EnterReturningClause is called when production returningClause is entered.
func (s *BasePartiQLListener) EnterReturningClause(ctx *ReturningClauseContext) {}

// ExitReturningClause is called when production returningClause is exited.
func (s *BasePartiQLListener) ExitReturningClause(ctx *ReturningClauseContext) {}

// EnterReturningColumn is called when production returningColumn is entered.
func (s *BasePartiQLListener) EnterReturningColumn(ctx *ReturningColumnContext) {}

// ExitReturningColumn is called when production returningColumn is exited.
func (s *BasePartiQLListener) ExitReturningColumn(ctx *ReturningColumnContext) {}

// EnterFromClauseSimpleExplicit is called when production FromClauseSimpleExplicit is entered.
func (s *BasePartiQLListener) EnterFromClauseSimpleExplicit(ctx *FromClauseSimpleExplicitContext) {}

// ExitFromClauseSimpleExplicit is called when production FromClauseSimpleExplicit is exited.
func (s *BasePartiQLListener) ExitFromClauseSimpleExplicit(ctx *FromClauseSimpleExplicitContext) {}

// EnterFromClauseSimpleImplicit is called when production FromClauseSimpleImplicit is entered.
func (s *BasePartiQLListener) EnterFromClauseSimpleImplicit(ctx *FromClauseSimpleImplicitContext) {}

// ExitFromClauseSimpleImplicit is called when production FromClauseSimpleImplicit is exited.
func (s *BasePartiQLListener) ExitFromClauseSimpleImplicit(ctx *FromClauseSimpleImplicitContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BasePartiQLListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BasePartiQLListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterSelectAll is called when production SelectAll is entered.
func (s *BasePartiQLListener) EnterSelectAll(ctx *SelectAllContext) {}

// ExitSelectAll is called when production SelectAll is exited.
func (s *BasePartiQLListener) ExitSelectAll(ctx *SelectAllContext) {}

// EnterSelectItems is called when production SelectItems is entered.
func (s *BasePartiQLListener) EnterSelectItems(ctx *SelectItemsContext) {}

// ExitSelectItems is called when production SelectItems is exited.
func (s *BasePartiQLListener) ExitSelectItems(ctx *SelectItemsContext) {}

// EnterSelectValue is called when production SelectValue is entered.
func (s *BasePartiQLListener) EnterSelectValue(ctx *SelectValueContext) {}

// ExitSelectValue is called when production SelectValue is exited.
func (s *BasePartiQLListener) ExitSelectValue(ctx *SelectValueContext) {}

// EnterSelectPivot is called when production SelectPivot is entered.
func (s *BasePartiQLListener) EnterSelectPivot(ctx *SelectPivotContext) {}

// ExitSelectPivot is called when production SelectPivot is exited.
func (s *BasePartiQLListener) ExitSelectPivot(ctx *SelectPivotContext) {}

// EnterProjectionItems is called when production projectionItems is entered.
func (s *BasePartiQLListener) EnterProjectionItems(ctx *ProjectionItemsContext) {}

// ExitProjectionItems is called when production projectionItems is exited.
func (s *BasePartiQLListener) ExitProjectionItems(ctx *ProjectionItemsContext) {}

// EnterProjectionItem is called when production projectionItem is entered.
func (s *BasePartiQLListener) EnterProjectionItem(ctx *ProjectionItemContext) {}

// ExitProjectionItem is called when production projectionItem is exited.
func (s *BasePartiQLListener) ExitProjectionItem(ctx *ProjectionItemContext) {}

// EnterSetQuantifierStrategy is called when production setQuantifierStrategy is entered.
func (s *BasePartiQLListener) EnterSetQuantifierStrategy(ctx *SetQuantifierStrategyContext) {}

// ExitSetQuantifierStrategy is called when production setQuantifierStrategy is exited.
func (s *BasePartiQLListener) ExitSetQuantifierStrategy(ctx *SetQuantifierStrategyContext) {}

// EnterLetClause is called when production letClause is entered.
func (s *BasePartiQLListener) EnterLetClause(ctx *LetClauseContext) {}

// ExitLetClause is called when production letClause is exited.
func (s *BasePartiQLListener) ExitLetClause(ctx *LetClauseContext) {}

// EnterLetBinding is called when production letBinding is entered.
func (s *BasePartiQLListener) EnterLetBinding(ctx *LetBindingContext) {}

// ExitLetBinding is called when production letBinding is exited.
func (s *BasePartiQLListener) ExitLetBinding(ctx *LetBindingContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BasePartiQLListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BasePartiQLListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderSortSpec is called when production orderSortSpec is entered.
func (s *BasePartiQLListener) EnterOrderSortSpec(ctx *OrderSortSpecContext) {}

// ExitOrderSortSpec is called when production orderSortSpec is exited.
func (s *BasePartiQLListener) ExitOrderSortSpec(ctx *OrderSortSpecContext) {}

// EnterGroupClause is called when production groupClause is entered.
func (s *BasePartiQLListener) EnterGroupClause(ctx *GroupClauseContext) {}

// ExitGroupClause is called when production groupClause is exited.
func (s *BasePartiQLListener) ExitGroupClause(ctx *GroupClauseContext) {}

// EnterGroupAlias is called when production groupAlias is entered.
func (s *BasePartiQLListener) EnterGroupAlias(ctx *GroupAliasContext) {}

// ExitGroupAlias is called when production groupAlias is exited.
func (s *BasePartiQLListener) ExitGroupAlias(ctx *GroupAliasContext) {}

// EnterGroupKey is called when production groupKey is entered.
func (s *BasePartiQLListener) EnterGroupKey(ctx *GroupKeyContext) {}

// ExitGroupKey is called when production groupKey is exited.
func (s *BasePartiQLListener) ExitGroupKey(ctx *GroupKeyContext) {}

// EnterOver is called when production over is entered.
func (s *BasePartiQLListener) EnterOver(ctx *OverContext) {}

// ExitOver is called when production over is exited.
func (s *BasePartiQLListener) ExitOver(ctx *OverContext) {}

// EnterWindowPartitionList is called when production windowPartitionList is entered.
func (s *BasePartiQLListener) EnterWindowPartitionList(ctx *WindowPartitionListContext) {}

// ExitWindowPartitionList is called when production windowPartitionList is exited.
func (s *BasePartiQLListener) ExitWindowPartitionList(ctx *WindowPartitionListContext) {}

// EnterWindowSortSpecList is called when production windowSortSpecList is entered.
func (s *BasePartiQLListener) EnterWindowSortSpecList(ctx *WindowSortSpecListContext) {}

// ExitWindowSortSpecList is called when production windowSortSpecList is exited.
func (s *BasePartiQLListener) ExitWindowSortSpecList(ctx *WindowSortSpecListContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BasePartiQLListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BasePartiQLListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BasePartiQLListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BasePartiQLListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterWhereClauseSelect is called when production whereClauseSelect is entered.
func (s *BasePartiQLListener) EnterWhereClauseSelect(ctx *WhereClauseSelectContext) {}

// ExitWhereClauseSelect is called when production whereClauseSelect is exited.
func (s *BasePartiQLListener) ExitWhereClauseSelect(ctx *WhereClauseSelectContext) {}

// EnterOffsetByClause is called when production offsetByClause is entered.
func (s *BasePartiQLListener) EnterOffsetByClause(ctx *OffsetByClauseContext) {}

// ExitOffsetByClause is called when production offsetByClause is exited.
func (s *BasePartiQLListener) ExitOffsetByClause(ctx *OffsetByClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BasePartiQLListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BasePartiQLListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterGpmlPattern is called when production gpmlPattern is entered.
func (s *BasePartiQLListener) EnterGpmlPattern(ctx *GpmlPatternContext) {}

// ExitGpmlPattern is called when production gpmlPattern is exited.
func (s *BasePartiQLListener) ExitGpmlPattern(ctx *GpmlPatternContext) {}

// EnterGpmlPatternList is called when production gpmlPatternList is entered.
func (s *BasePartiQLListener) EnterGpmlPatternList(ctx *GpmlPatternListContext) {}

// ExitGpmlPatternList is called when production gpmlPatternList is exited.
func (s *BasePartiQLListener) ExitGpmlPatternList(ctx *GpmlPatternListContext) {}

// EnterMatchPattern is called when production matchPattern is entered.
func (s *BasePartiQLListener) EnterMatchPattern(ctx *MatchPatternContext) {}

// ExitMatchPattern is called when production matchPattern is exited.
func (s *BasePartiQLListener) ExitMatchPattern(ctx *MatchPatternContext) {}

// EnterGraphPart is called when production graphPart is entered.
func (s *BasePartiQLListener) EnterGraphPart(ctx *GraphPartContext) {}

// ExitGraphPart is called when production graphPart is exited.
func (s *BasePartiQLListener) ExitGraphPart(ctx *GraphPartContext) {}

// EnterSelectorBasic is called when production SelectorBasic is entered.
func (s *BasePartiQLListener) EnterSelectorBasic(ctx *SelectorBasicContext) {}

// ExitSelectorBasic is called when production SelectorBasic is exited.
func (s *BasePartiQLListener) ExitSelectorBasic(ctx *SelectorBasicContext) {}

// EnterSelectorAny is called when production SelectorAny is entered.
func (s *BasePartiQLListener) EnterSelectorAny(ctx *SelectorAnyContext) {}

// ExitSelectorAny is called when production SelectorAny is exited.
func (s *BasePartiQLListener) ExitSelectorAny(ctx *SelectorAnyContext) {}

// EnterSelectorShortest is called when production SelectorShortest is entered.
func (s *BasePartiQLListener) EnterSelectorShortest(ctx *SelectorShortestContext) {}

// ExitSelectorShortest is called when production SelectorShortest is exited.
func (s *BasePartiQLListener) ExitSelectorShortest(ctx *SelectorShortestContext) {}

// EnterPatternPathVariable is called when production patternPathVariable is entered.
func (s *BasePartiQLListener) EnterPatternPathVariable(ctx *PatternPathVariableContext) {}

// ExitPatternPathVariable is called when production patternPathVariable is exited.
func (s *BasePartiQLListener) ExitPatternPathVariable(ctx *PatternPathVariableContext) {}

// EnterPatternRestrictor is called when production patternRestrictor is entered.
func (s *BasePartiQLListener) EnterPatternRestrictor(ctx *PatternRestrictorContext) {}

// ExitPatternRestrictor is called when production patternRestrictor is exited.
func (s *BasePartiQLListener) ExitPatternRestrictor(ctx *PatternRestrictorContext) {}

// EnterNode is called when production node is entered.
func (s *BasePartiQLListener) EnterNode(ctx *NodeContext) {}

// ExitNode is called when production node is exited.
func (s *BasePartiQLListener) ExitNode(ctx *NodeContext) {}

// EnterEdgeWithSpec is called when production EdgeWithSpec is entered.
func (s *BasePartiQLListener) EnterEdgeWithSpec(ctx *EdgeWithSpecContext) {}

// ExitEdgeWithSpec is called when production EdgeWithSpec is exited.
func (s *BasePartiQLListener) ExitEdgeWithSpec(ctx *EdgeWithSpecContext) {}

// EnterEdgeAbbreviated is called when production EdgeAbbreviated is entered.
func (s *BasePartiQLListener) EnterEdgeAbbreviated(ctx *EdgeAbbreviatedContext) {}

// ExitEdgeAbbreviated is called when production EdgeAbbreviated is exited.
func (s *BasePartiQLListener) ExitEdgeAbbreviated(ctx *EdgeAbbreviatedContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BasePartiQLListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BasePartiQLListener) ExitPattern(ctx *PatternContext) {}

// EnterPatternQuantifier is called when production patternQuantifier is entered.
func (s *BasePartiQLListener) EnterPatternQuantifier(ctx *PatternQuantifierContext) {}

// ExitPatternQuantifier is called when production patternQuantifier is exited.
func (s *BasePartiQLListener) ExitPatternQuantifier(ctx *PatternQuantifierContext) {}

// EnterEdgeSpecRight is called when production EdgeSpecRight is entered.
func (s *BasePartiQLListener) EnterEdgeSpecRight(ctx *EdgeSpecRightContext) {}

// ExitEdgeSpecRight is called when production EdgeSpecRight is exited.
func (s *BasePartiQLListener) ExitEdgeSpecRight(ctx *EdgeSpecRightContext) {}

// EnterEdgeSpecUndirected is called when production EdgeSpecUndirected is entered.
func (s *BasePartiQLListener) EnterEdgeSpecUndirected(ctx *EdgeSpecUndirectedContext) {}

// ExitEdgeSpecUndirected is called when production EdgeSpecUndirected is exited.
func (s *BasePartiQLListener) ExitEdgeSpecUndirected(ctx *EdgeSpecUndirectedContext) {}

// EnterEdgeSpecLeft is called when production EdgeSpecLeft is entered.
func (s *BasePartiQLListener) EnterEdgeSpecLeft(ctx *EdgeSpecLeftContext) {}

// ExitEdgeSpecLeft is called when production EdgeSpecLeft is exited.
func (s *BasePartiQLListener) ExitEdgeSpecLeft(ctx *EdgeSpecLeftContext) {}

// EnterEdgeSpecUndirectedRight is called when production EdgeSpecUndirectedRight is entered.
func (s *BasePartiQLListener) EnterEdgeSpecUndirectedRight(ctx *EdgeSpecUndirectedRightContext) {}

// ExitEdgeSpecUndirectedRight is called when production EdgeSpecUndirectedRight is exited.
func (s *BasePartiQLListener) ExitEdgeSpecUndirectedRight(ctx *EdgeSpecUndirectedRightContext) {}

// EnterEdgeSpecUndirectedLeft is called when production EdgeSpecUndirectedLeft is entered.
func (s *BasePartiQLListener) EnterEdgeSpecUndirectedLeft(ctx *EdgeSpecUndirectedLeftContext) {}

// ExitEdgeSpecUndirectedLeft is called when production EdgeSpecUndirectedLeft is exited.
func (s *BasePartiQLListener) ExitEdgeSpecUndirectedLeft(ctx *EdgeSpecUndirectedLeftContext) {}

// EnterEdgeSpecBidirectional is called when production EdgeSpecBidirectional is entered.
func (s *BasePartiQLListener) EnterEdgeSpecBidirectional(ctx *EdgeSpecBidirectionalContext) {}

// ExitEdgeSpecBidirectional is called when production EdgeSpecBidirectional is exited.
func (s *BasePartiQLListener) ExitEdgeSpecBidirectional(ctx *EdgeSpecBidirectionalContext) {}

// EnterEdgeSpecUndirectedBidirectional is called when production EdgeSpecUndirectedBidirectional is entered.
func (s *BasePartiQLListener) EnterEdgeSpecUndirectedBidirectional(ctx *EdgeSpecUndirectedBidirectionalContext) {
}

// ExitEdgeSpecUndirectedBidirectional is called when production EdgeSpecUndirectedBidirectional is exited.
func (s *BasePartiQLListener) ExitEdgeSpecUndirectedBidirectional(ctx *EdgeSpecUndirectedBidirectionalContext) {
}

// EnterEdgeSpec is called when production edgeSpec is entered.
func (s *BasePartiQLListener) EnterEdgeSpec(ctx *EdgeSpecContext) {}

// ExitEdgeSpec is called when production edgeSpec is exited.
func (s *BasePartiQLListener) ExitEdgeSpec(ctx *EdgeSpecContext) {}

// EnterPatternPartLabel is called when production patternPartLabel is entered.
func (s *BasePartiQLListener) EnterPatternPartLabel(ctx *PatternPartLabelContext) {}

// ExitPatternPartLabel is called when production patternPartLabel is exited.
func (s *BasePartiQLListener) ExitPatternPartLabel(ctx *PatternPartLabelContext) {}

// EnterEdgeAbbrev is called when production edgeAbbrev is entered.
func (s *BasePartiQLListener) EnterEdgeAbbrev(ctx *EdgeAbbrevContext) {}

// ExitEdgeAbbrev is called when production edgeAbbrev is exited.
func (s *BasePartiQLListener) ExitEdgeAbbrev(ctx *EdgeAbbrevContext) {}

// EnterTableWrapped is called when production TableWrapped is entered.
func (s *BasePartiQLListener) EnterTableWrapped(ctx *TableWrappedContext) {}

// ExitTableWrapped is called when production TableWrapped is exited.
func (s *BasePartiQLListener) ExitTableWrapped(ctx *TableWrappedContext) {}

// EnterTableCrossJoin is called when production TableCrossJoin is entered.
func (s *BasePartiQLListener) EnterTableCrossJoin(ctx *TableCrossJoinContext) {}

// ExitTableCrossJoin is called when production TableCrossJoin is exited.
func (s *BasePartiQLListener) ExitTableCrossJoin(ctx *TableCrossJoinContext) {}

// EnterTableQualifiedJoin is called when production TableQualifiedJoin is entered.
func (s *BasePartiQLListener) EnterTableQualifiedJoin(ctx *TableQualifiedJoinContext) {}

// ExitTableQualifiedJoin is called when production TableQualifiedJoin is exited.
func (s *BasePartiQLListener) ExitTableQualifiedJoin(ctx *TableQualifiedJoinContext) {}

// EnterTableRefBase is called when production TableRefBase is entered.
func (s *BasePartiQLListener) EnterTableRefBase(ctx *TableRefBaseContext) {}

// ExitTableRefBase is called when production TableRefBase is exited.
func (s *BasePartiQLListener) ExitTableRefBase(ctx *TableRefBaseContext) {}

// EnterTableNonJoin is called when production tableNonJoin is entered.
func (s *BasePartiQLListener) EnterTableNonJoin(ctx *TableNonJoinContext) {}

// ExitTableNonJoin is called when production tableNonJoin is exited.
func (s *BasePartiQLListener) ExitTableNonJoin(ctx *TableNonJoinContext) {}

// EnterTableBaseRefSymbol is called when production TableBaseRefSymbol is entered.
func (s *BasePartiQLListener) EnterTableBaseRefSymbol(ctx *TableBaseRefSymbolContext) {}

// ExitTableBaseRefSymbol is called when production TableBaseRefSymbol is exited.
func (s *BasePartiQLListener) ExitTableBaseRefSymbol(ctx *TableBaseRefSymbolContext) {}

// EnterTableBaseRefClauses is called when production TableBaseRefClauses is entered.
func (s *BasePartiQLListener) EnterTableBaseRefClauses(ctx *TableBaseRefClausesContext) {}

// ExitTableBaseRefClauses is called when production TableBaseRefClauses is exited.
func (s *BasePartiQLListener) ExitTableBaseRefClauses(ctx *TableBaseRefClausesContext) {}

// EnterTableBaseRefMatch is called when production TableBaseRefMatch is entered.
func (s *BasePartiQLListener) EnterTableBaseRefMatch(ctx *TableBaseRefMatchContext) {}

// ExitTableBaseRefMatch is called when production TableBaseRefMatch is exited.
func (s *BasePartiQLListener) ExitTableBaseRefMatch(ctx *TableBaseRefMatchContext) {}

// EnterTableUnpivot is called when production tableUnpivot is entered.
func (s *BasePartiQLListener) EnterTableUnpivot(ctx *TableUnpivotContext) {}

// ExitTableUnpivot is called when production tableUnpivot is exited.
func (s *BasePartiQLListener) ExitTableUnpivot(ctx *TableUnpivotContext) {}

// EnterJoinRhsBase is called when production JoinRhsBase is entered.
func (s *BasePartiQLListener) EnterJoinRhsBase(ctx *JoinRhsBaseContext) {}

// ExitJoinRhsBase is called when production JoinRhsBase is exited.
func (s *BasePartiQLListener) ExitJoinRhsBase(ctx *JoinRhsBaseContext) {}

// EnterJoinRhsTableJoined is called when production JoinRhsTableJoined is entered.
func (s *BasePartiQLListener) EnterJoinRhsTableJoined(ctx *JoinRhsTableJoinedContext) {}

// ExitJoinRhsTableJoined is called when production JoinRhsTableJoined is exited.
func (s *BasePartiQLListener) ExitJoinRhsTableJoined(ctx *JoinRhsTableJoinedContext) {}

// EnterJoinSpec is called when production joinSpec is entered.
func (s *BasePartiQLListener) EnterJoinSpec(ctx *JoinSpecContext) {}

// ExitJoinSpec is called when production joinSpec is exited.
func (s *BasePartiQLListener) ExitJoinSpec(ctx *JoinSpecContext) {}

// EnterJoinType is called when production joinType is entered.
func (s *BasePartiQLListener) EnterJoinType(ctx *JoinTypeContext) {}

// ExitJoinType is called when production joinType is exited.
func (s *BasePartiQLListener) ExitJoinType(ctx *JoinTypeContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePartiQLListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePartiQLListener) ExitExpr(ctx *ExprContext) {}

// EnterIntersect is called when production Intersect is entered.
func (s *BasePartiQLListener) EnterIntersect(ctx *IntersectContext) {}

// ExitIntersect is called when production Intersect is exited.
func (s *BasePartiQLListener) ExitIntersect(ctx *IntersectContext) {}

// EnterQueryBase is called when production QueryBase is entered.
func (s *BasePartiQLListener) EnterQueryBase(ctx *QueryBaseContext) {}

// ExitQueryBase is called when production QueryBase is exited.
func (s *BasePartiQLListener) ExitQueryBase(ctx *QueryBaseContext) {}

// EnterExcept is called when production Except is entered.
func (s *BasePartiQLListener) EnterExcept(ctx *ExceptContext) {}

// ExitExcept is called when production Except is exited.
func (s *BasePartiQLListener) ExitExcept(ctx *ExceptContext) {}

// EnterUnion is called when production Union is entered.
func (s *BasePartiQLListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production Union is exited.
func (s *BasePartiQLListener) ExitUnion(ctx *UnionContext) {}

// EnterSfwQuery is called when production SfwQuery is entered.
func (s *BasePartiQLListener) EnterSfwQuery(ctx *SfwQueryContext) {}

// ExitSfwQuery is called when production SfwQuery is exited.
func (s *BasePartiQLListener) ExitSfwQuery(ctx *SfwQueryContext) {}

// EnterSfwBase is called when production SfwBase is entered.
func (s *BasePartiQLListener) EnterSfwBase(ctx *SfwBaseContext) {}

// ExitSfwBase is called when production SfwBase is exited.
func (s *BasePartiQLListener) ExitSfwBase(ctx *SfwBaseContext) {}

// EnterOr is called when production Or is entered.
func (s *BasePartiQLListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BasePartiQLListener) ExitOr(ctx *OrContext) {}

// EnterExprOrBase is called when production ExprOrBase is entered.
func (s *BasePartiQLListener) EnterExprOrBase(ctx *ExprOrBaseContext) {}

// ExitExprOrBase is called when production ExprOrBase is exited.
func (s *BasePartiQLListener) ExitExprOrBase(ctx *ExprOrBaseContext) {}

// EnterExprAndBase is called when production ExprAndBase is entered.
func (s *BasePartiQLListener) EnterExprAndBase(ctx *ExprAndBaseContext) {}

// ExitExprAndBase is called when production ExprAndBase is exited.
func (s *BasePartiQLListener) ExitExprAndBase(ctx *ExprAndBaseContext) {}

// EnterAnd is called when production And is entered.
func (s *BasePartiQLListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BasePartiQLListener) ExitAnd(ctx *AndContext) {}

// EnterNot is called when production Not is entered.
func (s *BasePartiQLListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BasePartiQLListener) ExitNot(ctx *NotContext) {}

// EnterExprNotBase is called when production ExprNotBase is entered.
func (s *BasePartiQLListener) EnterExprNotBase(ctx *ExprNotBaseContext) {}

// ExitExprNotBase is called when production ExprNotBase is exited.
func (s *BasePartiQLListener) ExitExprNotBase(ctx *ExprNotBaseContext) {}

// EnterPredicateIn is called when production PredicateIn is entered.
func (s *BasePartiQLListener) EnterPredicateIn(ctx *PredicateInContext) {}

// ExitPredicateIn is called when production PredicateIn is exited.
func (s *BasePartiQLListener) ExitPredicateIn(ctx *PredicateInContext) {}

// EnterPredicateBetween is called when production PredicateBetween is entered.
func (s *BasePartiQLListener) EnterPredicateBetween(ctx *PredicateBetweenContext) {}

// ExitPredicateBetween is called when production PredicateBetween is exited.
func (s *BasePartiQLListener) ExitPredicateBetween(ctx *PredicateBetweenContext) {}

// EnterPredicateBase is called when production PredicateBase is entered.
func (s *BasePartiQLListener) EnterPredicateBase(ctx *PredicateBaseContext) {}

// ExitPredicateBase is called when production PredicateBase is exited.
func (s *BasePartiQLListener) ExitPredicateBase(ctx *PredicateBaseContext) {}

// EnterPredicateComparison is called when production PredicateComparison is entered.
func (s *BasePartiQLListener) EnterPredicateComparison(ctx *PredicateComparisonContext) {}

// ExitPredicateComparison is called when production PredicateComparison is exited.
func (s *BasePartiQLListener) ExitPredicateComparison(ctx *PredicateComparisonContext) {}

// EnterPredicateIs is called when production PredicateIs is entered.
func (s *BasePartiQLListener) EnterPredicateIs(ctx *PredicateIsContext) {}

// ExitPredicateIs is called when production PredicateIs is exited.
func (s *BasePartiQLListener) ExitPredicateIs(ctx *PredicateIsContext) {}

// EnterPredicateLike is called when production PredicateLike is entered.
func (s *BasePartiQLListener) EnterPredicateLike(ctx *PredicateLikeContext) {}

// ExitPredicateLike is called when production PredicateLike is exited.
func (s *BasePartiQLListener) ExitPredicateLike(ctx *PredicateLikeContext) {}

// EnterMathOp00 is called when production mathOp00 is entered.
func (s *BasePartiQLListener) EnterMathOp00(ctx *MathOp00Context) {}

// ExitMathOp00 is called when production mathOp00 is exited.
func (s *BasePartiQLListener) ExitMathOp00(ctx *MathOp00Context) {}

// EnterMathOp01 is called when production mathOp01 is entered.
func (s *BasePartiQLListener) EnterMathOp01(ctx *MathOp01Context) {}

// ExitMathOp01 is called when production mathOp01 is exited.
func (s *BasePartiQLListener) ExitMathOp01(ctx *MathOp01Context) {}

// EnterMathOp02 is called when production mathOp02 is entered.
func (s *BasePartiQLListener) EnterMathOp02(ctx *MathOp02Context) {}

// ExitMathOp02 is called when production mathOp02 is exited.
func (s *BasePartiQLListener) ExitMathOp02(ctx *MathOp02Context) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BasePartiQLListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BasePartiQLListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterExprPrimaryPath is called when production ExprPrimaryPath is entered.
func (s *BasePartiQLListener) EnterExprPrimaryPath(ctx *ExprPrimaryPathContext) {}

// ExitExprPrimaryPath is called when production ExprPrimaryPath is exited.
func (s *BasePartiQLListener) ExitExprPrimaryPath(ctx *ExprPrimaryPathContext) {}

// EnterExprPrimaryBase is called when production ExprPrimaryBase is entered.
func (s *BasePartiQLListener) EnterExprPrimaryBase(ctx *ExprPrimaryBaseContext) {}

// ExitExprPrimaryBase is called when production ExprPrimaryBase is exited.
func (s *BasePartiQLListener) ExitExprPrimaryBase(ctx *ExprPrimaryBaseContext) {}

// EnterExprTermWrappedQuery is called when production ExprTermWrappedQuery is entered.
func (s *BasePartiQLListener) EnterExprTermWrappedQuery(ctx *ExprTermWrappedQueryContext) {}

// ExitExprTermWrappedQuery is called when production ExprTermWrappedQuery is exited.
func (s *BasePartiQLListener) ExitExprTermWrappedQuery(ctx *ExprTermWrappedQueryContext) {}

// EnterExprTermBase is called when production ExprTermBase is entered.
func (s *BasePartiQLListener) EnterExprTermBase(ctx *ExprTermBaseContext) {}

// ExitExprTermBase is called when production ExprTermBase is exited.
func (s *BasePartiQLListener) ExitExprTermBase(ctx *ExprTermBaseContext) {}

// EnterNullIf is called when production nullIf is entered.
func (s *BasePartiQLListener) EnterNullIf(ctx *NullIfContext) {}

// ExitNullIf is called when production nullIf is exited.
func (s *BasePartiQLListener) ExitNullIf(ctx *NullIfContext) {}

// EnterCoalesce is called when production coalesce is entered.
func (s *BasePartiQLListener) EnterCoalesce(ctx *CoalesceContext) {}

// ExitCoalesce is called when production coalesce is exited.
func (s *BasePartiQLListener) ExitCoalesce(ctx *CoalesceContext) {}

// EnterCaseExpr is called when production caseExpr is entered.
func (s *BasePartiQLListener) EnterCaseExpr(ctx *CaseExprContext) {}

// ExitCaseExpr is called when production caseExpr is exited.
func (s *BasePartiQLListener) ExitCaseExpr(ctx *CaseExprContext) {}

// EnterValues is called when production values is entered.
func (s *BasePartiQLListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production values is exited.
func (s *BasePartiQLListener) ExitValues(ctx *ValuesContext) {}

// EnterValueRow is called when production valueRow is entered.
func (s *BasePartiQLListener) EnterValueRow(ctx *ValueRowContext) {}

// ExitValueRow is called when production valueRow is exited.
func (s *BasePartiQLListener) ExitValueRow(ctx *ValueRowContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BasePartiQLListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BasePartiQLListener) ExitValueList(ctx *ValueListContext) {}

// EnterSequenceConstructor is called when production sequenceConstructor is entered.
func (s *BasePartiQLListener) EnterSequenceConstructor(ctx *SequenceConstructorContext) {}

// ExitSequenceConstructor is called when production sequenceConstructor is exited.
func (s *BasePartiQLListener) ExitSequenceConstructor(ctx *SequenceConstructorContext) {}

// EnterSubstring is called when production substring is entered.
func (s *BasePartiQLListener) EnterSubstring(ctx *SubstringContext) {}

// ExitSubstring is called when production substring is exited.
func (s *BasePartiQLListener) ExitSubstring(ctx *SubstringContext) {}

// EnterCountAll is called when production CountAll is entered.
func (s *BasePartiQLListener) EnterCountAll(ctx *CountAllContext) {}

// ExitCountAll is called when production CountAll is exited.
func (s *BasePartiQLListener) ExitCountAll(ctx *CountAllContext) {}

// EnterAggregateBase is called when production AggregateBase is entered.
func (s *BasePartiQLListener) EnterAggregateBase(ctx *AggregateBaseContext) {}

// ExitAggregateBase is called when production AggregateBase is exited.
func (s *BasePartiQLListener) ExitAggregateBase(ctx *AggregateBaseContext) {}

// EnterLagLeadFunction is called when production LagLeadFunction is entered.
func (s *BasePartiQLListener) EnterLagLeadFunction(ctx *LagLeadFunctionContext) {}

// ExitLagLeadFunction is called when production LagLeadFunction is exited.
func (s *BasePartiQLListener) ExitLagLeadFunction(ctx *LagLeadFunctionContext) {}

// EnterCast is called when production cast is entered.
func (s *BasePartiQLListener) EnterCast(ctx *CastContext) {}

// ExitCast is called when production cast is exited.
func (s *BasePartiQLListener) ExitCast(ctx *CastContext) {}

// EnterCanLosslessCast is called when production canLosslessCast is entered.
func (s *BasePartiQLListener) EnterCanLosslessCast(ctx *CanLosslessCastContext) {}

// ExitCanLosslessCast is called when production canLosslessCast is exited.
func (s *BasePartiQLListener) ExitCanLosslessCast(ctx *CanLosslessCastContext) {}

// EnterCanCast is called when production canCast is entered.
func (s *BasePartiQLListener) EnterCanCast(ctx *CanCastContext) {}

// ExitCanCast is called when production canCast is exited.
func (s *BasePartiQLListener) ExitCanCast(ctx *CanCastContext) {}

// EnterExtract is called when production extract is entered.
func (s *BasePartiQLListener) EnterExtract(ctx *ExtractContext) {}

// ExitExtract is called when production extract is exited.
func (s *BasePartiQLListener) ExitExtract(ctx *ExtractContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BasePartiQLListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BasePartiQLListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterDateFunction is called when production dateFunction is entered.
func (s *BasePartiQLListener) EnterDateFunction(ctx *DateFunctionContext) {}

// ExitDateFunction is called when production dateFunction is exited.
func (s *BasePartiQLListener) ExitDateFunction(ctx *DateFunctionContext) {}

// EnterFunctionCallReserved is called when production FunctionCallReserved is entered.
func (s *BasePartiQLListener) EnterFunctionCallReserved(ctx *FunctionCallReservedContext) {}

// ExitFunctionCallReserved is called when production FunctionCallReserved is exited.
func (s *BasePartiQLListener) ExitFunctionCallReserved(ctx *FunctionCallReservedContext) {}

// EnterFunctionCallIdent is called when production FunctionCallIdent is entered.
func (s *BasePartiQLListener) EnterFunctionCallIdent(ctx *FunctionCallIdentContext) {}

// ExitFunctionCallIdent is called when production FunctionCallIdent is exited.
func (s *BasePartiQLListener) ExitFunctionCallIdent(ctx *FunctionCallIdentContext) {}

// EnterPathStepIndexExpr is called when production PathStepIndexExpr is entered.
func (s *BasePartiQLListener) EnterPathStepIndexExpr(ctx *PathStepIndexExprContext) {}

// ExitPathStepIndexExpr is called when production PathStepIndexExpr is exited.
func (s *BasePartiQLListener) ExitPathStepIndexExpr(ctx *PathStepIndexExprContext) {}

// EnterPathStepIndexAll is called when production PathStepIndexAll is entered.
func (s *BasePartiQLListener) EnterPathStepIndexAll(ctx *PathStepIndexAllContext) {}

// ExitPathStepIndexAll is called when production PathStepIndexAll is exited.
func (s *BasePartiQLListener) ExitPathStepIndexAll(ctx *PathStepIndexAllContext) {}

// EnterPathStepDotExpr is called when production PathStepDotExpr is entered.
func (s *BasePartiQLListener) EnterPathStepDotExpr(ctx *PathStepDotExprContext) {}

// ExitPathStepDotExpr is called when production PathStepDotExpr is exited.
func (s *BasePartiQLListener) ExitPathStepDotExpr(ctx *PathStepDotExprContext) {}

// EnterPathStepDotAll is called when production PathStepDotAll is entered.
func (s *BasePartiQLListener) EnterPathStepDotAll(ctx *PathStepDotAllContext) {}

// ExitPathStepDotAll is called when production PathStepDotAll is exited.
func (s *BasePartiQLListener) ExitPathStepDotAll(ctx *PathStepDotAllContext) {}

// EnterExprGraphMatchMany is called when production exprGraphMatchMany is entered.
func (s *BasePartiQLListener) EnterExprGraphMatchMany(ctx *ExprGraphMatchManyContext) {}

// ExitExprGraphMatchMany is called when production exprGraphMatchMany is exited.
func (s *BasePartiQLListener) ExitExprGraphMatchMany(ctx *ExprGraphMatchManyContext) {}

// EnterExprGraphMatchOne is called when production exprGraphMatchOne is entered.
func (s *BasePartiQLListener) EnterExprGraphMatchOne(ctx *ExprGraphMatchOneContext) {}

// ExitExprGraphMatchOne is called when production exprGraphMatchOne is exited.
func (s *BasePartiQLListener) ExitExprGraphMatchOne(ctx *ExprGraphMatchOneContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePartiQLListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePartiQLListener) ExitParameter(ctx *ParameterContext) {}

// EnterVarRefExpr is called when production varRefExpr is entered.
func (s *BasePartiQLListener) EnterVarRefExpr(ctx *VarRefExprContext) {}

// ExitVarRefExpr is called when production varRefExpr is exited.
func (s *BasePartiQLListener) ExitVarRefExpr(ctx *VarRefExprContext) {}

// EnterCollection is called when production collection is entered.
func (s *BasePartiQLListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BasePartiQLListener) ExitCollection(ctx *CollectionContext) {}

// EnterArray is called when production array is entered.
func (s *BasePartiQLListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BasePartiQLListener) ExitArray(ctx *ArrayContext) {}

// EnterBag is called when production bag is entered.
func (s *BasePartiQLListener) EnterBag(ctx *BagContext) {}

// ExitBag is called when production bag is exited.
func (s *BasePartiQLListener) ExitBag(ctx *BagContext) {}

// EnterTuple is called when production tuple is entered.
func (s *BasePartiQLListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production tuple is exited.
func (s *BasePartiQLListener) ExitTuple(ctx *TupleContext) {}

// EnterPair is called when production pair is entered.
func (s *BasePartiQLListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BasePartiQLListener) ExitPair(ctx *PairContext) {}

// EnterLiteralNull is called when production LiteralNull is entered.
func (s *BasePartiQLListener) EnterLiteralNull(ctx *LiteralNullContext) {}

// ExitLiteralNull is called when production LiteralNull is exited.
func (s *BasePartiQLListener) ExitLiteralNull(ctx *LiteralNullContext) {}

// EnterLiteralMissing is called when production LiteralMissing is entered.
func (s *BasePartiQLListener) EnterLiteralMissing(ctx *LiteralMissingContext) {}

// ExitLiteralMissing is called when production LiteralMissing is exited.
func (s *BasePartiQLListener) ExitLiteralMissing(ctx *LiteralMissingContext) {}

// EnterLiteralTrue is called when production LiteralTrue is entered.
func (s *BasePartiQLListener) EnterLiteralTrue(ctx *LiteralTrueContext) {}

// ExitLiteralTrue is called when production LiteralTrue is exited.
func (s *BasePartiQLListener) ExitLiteralTrue(ctx *LiteralTrueContext) {}

// EnterLiteralFalse is called when production LiteralFalse is entered.
func (s *BasePartiQLListener) EnterLiteralFalse(ctx *LiteralFalseContext) {}

// ExitLiteralFalse is called when production LiteralFalse is exited.
func (s *BasePartiQLListener) ExitLiteralFalse(ctx *LiteralFalseContext) {}

// EnterLiteralString is called when production LiteralString is entered.
func (s *BasePartiQLListener) EnterLiteralString(ctx *LiteralStringContext) {}

// ExitLiteralString is called when production LiteralString is exited.
func (s *BasePartiQLListener) ExitLiteralString(ctx *LiteralStringContext) {}

// EnterLiteralInteger is called when production LiteralInteger is entered.
func (s *BasePartiQLListener) EnterLiteralInteger(ctx *LiteralIntegerContext) {}

// ExitLiteralInteger is called when production LiteralInteger is exited.
func (s *BasePartiQLListener) ExitLiteralInteger(ctx *LiteralIntegerContext) {}

// EnterLiteralDecimal is called when production LiteralDecimal is entered.
func (s *BasePartiQLListener) EnterLiteralDecimal(ctx *LiteralDecimalContext) {}

// ExitLiteralDecimal is called when production LiteralDecimal is exited.
func (s *BasePartiQLListener) ExitLiteralDecimal(ctx *LiteralDecimalContext) {}

// EnterLiteralIon is called when production LiteralIon is entered.
func (s *BasePartiQLListener) EnterLiteralIon(ctx *LiteralIonContext) {}

// ExitLiteralIon is called when production LiteralIon is exited.
func (s *BasePartiQLListener) ExitLiteralIon(ctx *LiteralIonContext) {}

// EnterLiteralDate is called when production LiteralDate is entered.
func (s *BasePartiQLListener) EnterLiteralDate(ctx *LiteralDateContext) {}

// ExitLiteralDate is called when production LiteralDate is exited.
func (s *BasePartiQLListener) ExitLiteralDate(ctx *LiteralDateContext) {}

// EnterLiteralTime is called when production LiteralTime is entered.
func (s *BasePartiQLListener) EnterLiteralTime(ctx *LiteralTimeContext) {}

// ExitLiteralTime is called when production LiteralTime is exited.
func (s *BasePartiQLListener) ExitLiteralTime(ctx *LiteralTimeContext) {}

// EnterTypeAtomic is called when production TypeAtomic is entered.
func (s *BasePartiQLListener) EnterTypeAtomic(ctx *TypeAtomicContext) {}

// ExitTypeAtomic is called when production TypeAtomic is exited.
func (s *BasePartiQLListener) ExitTypeAtomic(ctx *TypeAtomicContext) {}

// EnterTypeArgSingle is called when production TypeArgSingle is entered.
func (s *BasePartiQLListener) EnterTypeArgSingle(ctx *TypeArgSingleContext) {}

// ExitTypeArgSingle is called when production TypeArgSingle is exited.
func (s *BasePartiQLListener) ExitTypeArgSingle(ctx *TypeArgSingleContext) {}

// EnterTypeVarChar is called when production TypeVarChar is entered.
func (s *BasePartiQLListener) EnterTypeVarChar(ctx *TypeVarCharContext) {}

// ExitTypeVarChar is called when production TypeVarChar is exited.
func (s *BasePartiQLListener) ExitTypeVarChar(ctx *TypeVarCharContext) {}

// EnterTypeArgDouble is called when production TypeArgDouble is entered.
func (s *BasePartiQLListener) EnterTypeArgDouble(ctx *TypeArgDoubleContext) {}

// ExitTypeArgDouble is called when production TypeArgDouble is exited.
func (s *BasePartiQLListener) ExitTypeArgDouble(ctx *TypeArgDoubleContext) {}

// EnterTypeTimeZone is called when production TypeTimeZone is entered.
func (s *BasePartiQLListener) EnterTypeTimeZone(ctx *TypeTimeZoneContext) {}

// ExitTypeTimeZone is called when production TypeTimeZone is exited.
func (s *BasePartiQLListener) ExitTypeTimeZone(ctx *TypeTimeZoneContext) {}

// EnterTypeCustom is called when production TypeCustom is entered.
func (s *BasePartiQLListener) EnterTypeCustom(ctx *TypeCustomContext) {}

// ExitTypeCustom is called when production TypeCustom is exited.
func (s *BasePartiQLListener) ExitTypeCustom(ctx *TypeCustomContext) {}
