package handler

type Input struct {
	// @inject_tag: db:"prev_addresses" json:"prev_addresses"
	PrevAddresses []string `protobuf:"bytes,1,rep,name=prev_addresses,proto3" json:"prev_addresses" db:"prev_addresses"`
	// @inject_tag: json:"prev_address_tag"
	PrevAddressTag string `protobuf:"bytes,2,opt,name=prev_address_tag,proto3" json:"prev_address_tag"`
	// @inject_tag: db:"prev_position" json:"prev_position"
	PrevPosition string `protobuf:"bytes,3,opt,name=prev_position,proto3" json:"prev_position" db:"prev_position"`
	// @inject_tag: db:"prev_tx_hash" json:"prev_tx_hash"
	PrevTxHash string `protobuf:"bytes,4,opt,name=prev_tx_hash,proto3" json:"prev_tx_hash" db:"prev_tx_hash"`
	// @inject_tag: db:"prev_type" json:"prev_type"
	PrevType string `protobuf:"bytes,5,opt,name=prev_type,proto3" json:"prev_type" db:"prev_type"`
	// @inject_tag: db:"prev_value" json:"prev_volume"
	PrevVolume string `protobuf:"bytes,6,opt,name=prev_volume,proto3" json:"prev_volume" db:"prev_value"`
	// @inject_tag: db:"script_asm" json:"script_asm"
	ScriptAsm string `protobuf:"bytes,7,opt,name=script_asm,proto3" json:"script_asm" db:"script_asm"`
	// @inject_tag: db:"script_hex" json:"script_hex"
	ScriptHex string `protobuf:"bytes,8,opt,name=script_hex,proto3" json:"script_hex" db:"script_hex"`
	// @inject_tag: json:"witness"
	Witness []string `protobuf:"bytes,9,rep,name=witness,proto3" json:"witness"`
}

type Output struct {
	// @inject_tag: db:"addresses" json:"addresses"
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses" db:"addresses"`
	// @inject_tag: json:"address_tag"
	AddressTag string `protobuf:"bytes,2,opt,name=address_tag,proto3" json:"address_tag"`
	// @inject_tag: db:"value" json:"volume"
	Volume string `protobuf:"bytes,3,opt,name=volume,proto3" json:"volume" db:"value"`
	// @inject_tag: db:"type" json:"type"
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type" db:"type"`
	// @inject_tag: db:"script_asm" json:"script_asm"
	ScriptAsm string `protobuf:"bytes,5,opt,name=script_asm,proto3" json:"script_asm" db:"script_asm"`
	// @inject_tag: db:"script_hex" json:"script_hex"
	ScriptHex string `protobuf:"bytes,6,opt,name=script_hex,proto3" json:"script_hex" db:"script_hex"`
	// @inject_tag: db:"spent_by_tx" json:"spent_by_tx"
	SpentByTx string `protobuf:"bytes,7,opt,name=spent_by_tx,proto3" json:"spent_by_tx" db:"spent_by_tx"`
	// @inject_tag: db:"spent_by_tx_position" json:"spent_by_tx_position"
	SpentByTxPosition string `protobuf:"bytes,8,opt,name=spent_by_tx_position,proto3" json:"spent_by_tx_position" db:"spent_by_tx_position"`
}

type Transaction struct {
	// @inject_tag: db:"height" json:"height"
	Height        string `protobuf:"bytes,1,opt,name=height,proto3" json:"height" db:"height"`
	Confirmations string `protobuf:"bytes,2,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	// @inject_tag: db:"size"
	Size string `protobuf:"bytes,3,opt,name=size,proto3" json:"size,omitempty" db:"size"`
	// @inject_tag: db:"vsize"
	Vsize string `protobuf:"bytes,4,opt,name=vsize,proto3" json:"vsize,omitempty" db:"vsize"`
	// @inject_tag: db:"weight"
	Weight string `protobuf:"bytes,5,opt,name=weight,proto3" json:"weight,omitempty" db:"weight"`
	// @inject_tag: db:"timestamp" json:"timestamp"
	Timestamp string `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp" db:"timestamp"`
	// @inject_tag: db:"fee" json:"fee"
	Fee     string `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee" db:"fee"`
	FeeRate string `protobuf:"bytes,8,opt,name=fee_rate,proto3" json:"fee_rate,omitempty"`
	// @inject_tag: db:"sigops"
	Sigops string `protobuf:"bytes,9,opt,name=sigops,proto3" json:"sigops,omitempty" db:"sigops"`
	// @inject_tag: db:"inputs_count" json:"inputs_count"
	InputsCount string `protobuf:"bytes,10,opt,name=inputs_count,proto3" json:"inputs_count" db:"inputs_count"`
	// @inject_tag: db:"inputs_value" json:"inputs_volume"
	InputsVolume string `protobuf:"bytes,11,opt,name=inputs_volume,proto3" json:"inputs_volume" db:"inputs_value"`
	// @inject_tag: db:"outputs_count" json:"outputs_count"
	OutputsCount string `protobuf:"bytes,12,opt,name=outputs_count,proto3" json:"outputs_count" db:"outputs_count"`
	// @inject_tag: db:"outputs_value" json:"outputs_volume"
	OutputsVolume string `protobuf:"bytes,13,opt,name=outputs_volume,proto3" json:"outputs_volume" db:"outputs_value"`
	// @inject_tag: db:"hash" json:"hash"
	Hash    string `protobuf:"bytes,14,opt,name=hash,proto3" json:"hash" db:"hash"`
	Witness string `protobuf:"bytes,15,opt,name=witness,proto3" json:"witness,omitempty"`
	// @inject_tag: db:"is_coinbase" json:"is_coinbase"
	IsCoinbase string `protobuf:"bytes,16,opt,name=is_coinbase,proto3" json:"is_coinbase" db:"is_coinbase"`
	Status     string `protobuf:"bytes,17,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: json:"inputs"
	Inputs []*Input `protobuf:"bytes,18,rep,name=inputs,proto3" json:"inputs"`
	// @inject_tag: json:"outputs"
	Outputs     []*Output `protobuf:"bytes,19,rep,name=outputs,proto3" json:"outputs"`
	BalanceDiff string    `protobuf:"bytes,20,opt,name=balance_diff,proto3" json:"balance_diff,omitempty"`
	Coin        string    `protobuf:"bytes,21,opt,name=coin,proto3" json:"coin,omitempty"`
}
