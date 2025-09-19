CREATE TYPE org_memberships_status AS ENUM ('active', 'suspended', 'ended', 'pending', 'revoked');

CREATE TABLE IF NOT EXISTS org_memberships (
    id UUID PRIMARY KEY,
    org_id UUID REFERENCES organizations(id) NOT NULL,
    user_id UUID REFERENCES users(id) NOT NULL,
    role_id UUID REFERENCES roles(id) NOT NULL,
    status org_memberships_status NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    finalized_by UUID REFERENCES users(id) NULL,
    finalized_at TIMESTAMP WITH TIME ZONE NULL

)

CREATE INDEX idx_org_memberships_user_id ON org_memberships(user_id);
CREATE INDEX idx_org_memberships_org_id ON org_memberships(org_id);
