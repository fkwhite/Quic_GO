package mocks

//go:generate sh -c "mockgen -package mockquic -destination quic/stream.go github.com/fkwhite/Quic_GO Stream"
//go:generate sh -c "mockgen -package mockquic -destination quic/early_session_tmp.go github.com/fkwhite/Quic_GO EarlySession && sed 's/qtls.ConnectionState/quic.ConnectionState/g' quic/early_session_tmp.go > quic/early_session.go && rm quic/early_session_tmp.go && goimports -w quic/early_session.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/early_listener.go github.com/fkwhite/Quic_GO EarlyListener"
//go:generate sh -c "mockgen -package mocklogging -destination logging/tracer.go github.com/fkwhite/Quic_GO/logging Tracer"
//go:generate sh -c "mockgen -package mocklogging -destination logging/connection_tracer.go github.com/fkwhite/Quic_GO/logging ConnectionTracer"
//go:generate sh -c "mockgen -package mocks -destination short_header_sealer.go github.com/fkwhite/Quic_GO/internal/handshake ShortHeaderSealer"
//go:generate sh -c "mockgen -package mocks -destination short_header_opener.go github.com/fkwhite/Quic_GO/internal/handshake ShortHeaderOpener"
//go:generate sh -c "mockgen -package mocks -destination long_header_opener.go github.com/fkwhite/Quic_GO/internal/handshake LongHeaderOpener"
//go:generate sh -c "mockgen -package mocks -destination crypto_setup_tmp.go github.com/fkwhite/Quic_GO/internal/handshake CryptoSetup && sed -E 's~github.com/marten-seemann/qtls[[:alnum:]_-]*~github.com/fkwhite/Quic_GO/internal/qtls~g; s~qtls.ConnectionStateWith0RTT~qtls.ConnectionState~g' crypto_setup_tmp.go > crypto_setup.go && rm crypto_setup_tmp.go && goimports -w crypto_setup.go"
//go:generate sh -c "mockgen -package mocks -destination stream_flow_controller.go github.com/fkwhite/Quic_GO/internal/flowcontrol StreamFlowController"
//go:generate sh -c "mockgen -package mocks -destination congestion.go github.com/fkwhite/Quic_GO/internal/congestion SendAlgorithmWithDebugInfos"
//go:generate sh -c "mockgen -package mocks -destination connection_flow_controller.go github.com/fkwhite/Quic_GO/internal/flowcontrol ConnectionFlowController"
//go:generate sh -c "mockgen -package mockackhandler -destination ackhandler/sent_packet_handler.go github.com/fkwhite/Quic_GO/internal/ackhandler SentPacketHandler"
//go:generate sh -c "mockgen -package mockackhandler -destination ackhandler/received_packet_handler.go github.com/fkwhite/Quic_GO/internal/ackhandler ReceivedPacketHandler"

// The following command produces a warning message on OSX, however, it still generates the correct mock file.
// See https://github.com/golang/mock/issues/339 for details.
//go:generate sh -c "mockgen -package mocktls -destination tls/client_session_cache.go crypto/tls ClientSessionCache"
